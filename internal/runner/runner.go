package runner

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oneaudit/nuclei-ng/internal/nuclei"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	"github.com/oneaudit/nuclei-ng/pkg/utils/extensions"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
	errorutil "github.com/projectdiscovery/utils/errors"
	"gopkg.in/yaml.v3"
	"math"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func Execute(options *types.Options) error {
	options.ConfigureOutput()
	showBanner()

	if options.Version {
		gologger.Info().Msgf("Current version: %s", version)
		return nil
	}

	if err := validateOptions(options); err != nil {
		return errorutil.NewWithErr(err).Msgf("could not validate options")
	}

	spec, err := openapi3.NewLoader().LoadFromFile(options.InputFile)
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("Error loading Swagger file")
	}
	var entries = make(map[string]*openapi3.Paths)
	entries["generic"] = &openapi3.Paths{}
	entries["images"] = &openapi3.Paths{}
	entries["html"] = &openapi3.Paths{}

	for path, item := range spec.Paths.Map() {
		entries["generic"].Set(path, item)
		ext := filepath.Ext(path)

		// HTML
		if ext == "" {
			entries["html"].Set(path, item)
		} else {
			found := false
			for _, htmlExt := range extensions.DefaultDenylist {
				if strings.EqualFold(ext, htmlExt) {
					found = true
					break
				}
			}
			if !found {
				entries["html"].Set(path, item)
			}
		}
	}

	writer := nuclei.NewNGStandardWriter()

	for category, paths := range entries {
		if paths.Len() == 0 {
			continue
		}

		// Compute nuclei options
		var tags []string
		switch category {
		case "image":
			tags = append(tags, "image")
		case "generic":
			tags = append(tags, "generic")
		case "html":
			tags = append(tags, "html")
		}
		gologger.Info().Msgf("Running nuclei with tags: %v against %d targets\n", tags, paths.Len())

		// Create a temporary swagger file
		// With the paths associated to the tag
		tempFile, err := os.CreateTemp("", "swagger.yaml")
		if err != nil {
			return errorutil.NewWithErr(err).Msgf("Error creating temp file")
		}
		//goland:noinspection GoDeferInLoop,GoUnhandledErrorResult
		defer tempFile.Close()
		gologger.Info().Msgf("Temporary file created: %s\n", tempFile.Name())
		encoder := yaml.NewEncoder(tempFile)
		encoder.SetIndent(2)
		spec.Paths = paths
		err = encoder.Encode(&spec)
		if err != nil {
			return errorutil.NewWithErr(err).Msgf("could not write output file: %s", tempFile.Name())
		}
		//goland:noinspection GoDeferInLoop,GoUnhandledErrorResult
		defer os.Remove(tempFile.Name())

		nucleiTemplateDir := "../nuclei-dast-templates/"
		//nucleiTemplateDir := "../nuclei-dast-templates/web/4.fingerprint/cookies/"
		//nucleiTemplateDir := "../nuclei-dast-templates/web/4.fingerprint/pages/"

		// Run nuclei
		cmd := exec.Command("nuclei",
			"-dast", "-no-interactsh",

			"-im", "openapi",
			"-list", tempFile.Name(),

			"-disable-update-check",
			"-t", nucleiTemplateDir,
			"-update-template-dir", nucleiTemplateDir,

			"-tags", strings.Join(tags, " "),

			"-follow-redirects", "-no-mhe",
			"-fuzz-param-frequency", strconv.Itoa(math.MaxInt32-1),

			"-j", "-silent", "-omit-raw", "-omit-template", "-no-color",
		)
		cmdOutput, err := cmd.CombinedOutput()
		if err != nil {
			return errorutil.NewWithErr(err).Msgf("Error executing nuclei command: %s", cmdOutput)
		}
		lines := strings.Split(string(cmdOutput), "\n")

		countMap := make(map[string]int)
		dataMap := make(map[string]*output.ResultEvent)
		endpointsMap := make(map[string][]string)
		for _, line := range lines {
			if line == "" {
				continue
			}
			var result output.ResultEvent
			err := json.Unmarshal([]byte(line), &result)
			if err != nil {
				gologger.Error().Msgf("Error unmarshaling JSON: %v for %s", err.Error(), line)
				continue
			}

			// Create the key in the format [template-id:matcher-name]
			extracted := result.ExtractedResults
			if len(extracted) == 0 {
				extracted = append(extracted, "")
			}

			for _, v := range extracted {
				key := fmt.Sprintf("[%s:%s:%s]", result.TemplateID, result.MatcherName, v)
				countMap[key]++
				dataMap[key] = &result
				if _, ok := endpointsMap[key]; !ok {
					endpointsMap[key] = []string{}
				}
				if result.Matched != "" {
					parsedUrl, err := url.Parse(result.Matched)
					var newValue string
					if err == nil {
						newValue = parsedUrl.Path
					} else {
						newValue = result.Matched
					}
					if newValue == "" {
						continue
					}
					found := false
					for _, endpoint := range endpointsMap[key] {
						if endpoint == newValue {
							found = true
							break
						}
					}
					if !found {
						endpointsMap[key] = append(endpointsMap[key], newValue)
					}
				}
			}
		}

		keys := make([]string, 0, len(countMap))
		for key := range countMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			result := dataMap[key]
			count := countMap[key]
			// Add a teaser
			sort.Strings(endpointsMap[key])
			if count > 3 {
				endpointsMap[key] = endpointsMap[key][:3]
				endpointsMap[key] = append(endpointsMap[key], "...")
			}
			// Change message shown
			result.Matched = ""
			result.Host = fmt.Sprintf("Found on %d URLs", count)
			// Do not display fake fuzzing information
			if !strings.Contains(result.Info.Tags.String(), "fuzzing") {
				result.IsFuzzingResult = false
			}
			// Print to stdout
			_, _ = os.Stdout.Write(writer.FormatScreen(result, endpointsMap[key]))
			_, _ = os.Stdout.Write([]byte("\n"))
		}
	}

	return nil
}
