package nuclei

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	openapiutil "github.com/oneaudit/nuclei-ng/pkg/utils/openapi"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
	errorutil "github.com/projectdiscovery/utils/errors"
	"math"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ExecuteCommand(options *types.Options, tags types.Tag, specification *openapi3.T, paths *openapi3.Paths) (string, error) {
	args := []string{
		"-dast", "-no-interactsh",

		"-disable-update-check",
		"-update-template-dir", options.NucleiTemplateDir,

		"-follow-redirects", "-no-mhe",
		"-fuzz-param-frequency", strconv.Itoa(math.MaxInt32 - 1),

		"-j", "-silent", "-omit-raw", "-omit-template", "-no-color",

		"-config", options.NucleiConfig,
	}

	tempFile, err := openapiutil.CreateTemporarySwaggerFile(specification, paths)
	if err != nil {
		return "", err
	}

	args = append(args, "-im", "openapi", "-list", tempFile.Name())

	if tags == types.WordPress {
		// Load a specific workflow that will gradually enable tags
		args = append(args, "-w", "workflows/wordpress.yaml")
	} else {
		// Load all templates while filtering them using tags
		args = append(args,
			"-t", options.NucleiTemplateDir,
			"-tags", string(tags),
			//"-etags", "noisy,fuzzing",
		)
	}

	//goland:noinspection GoDeferInLoop,GoUnhandledErrorResult
	defer os.Remove(tempFile.Name())

	// Run the command
	cmd := exec.Command("nuclei", args...)
	gologger.Debug().Msgf("Executing command: %v", cmd)
	//return "", errorutil.New("Executing command: %v", cmd)

	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		return "", errorutil.NewWithErr(err).Msgf("%s", cmdOutput)
	}

	return string(cmdOutput), nil
}

type ParsedEvent struct {
	Name      string             `json:"name"`
	Result    output.ResultEvent `json:"result"`
	Endpoints []string           `json:"endpoints"`
	Count     int                `json:"count"`
}

func ParseResult(result string) map[string]*ParsedEvent {
	endpointsMap := make(map[string]*ParsedEvent)

	lines := strings.Split(result, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Parse Line
		var result output.ResultEvent
		err := json.Unmarshal([]byte(line), &result)
		if err != nil {
			gologger.Error().Msgf("Error unmarshaling JSON: %v for %s", err.Error(), line)
			continue
		}

		// We want to have one extracted result per entry
		extracted := result.ExtractedResults
		if len(extracted) == 0 {
			extracted = append(extracted, "")
		}

		// Nuclei JavaScript templates are not returning the path
		// So, we have to add trick like these
		var extractedCleaned []string
		var baseEndpoints []string
		for _, extractedResult := range extracted {
			if strings.HasPrefix(extractedResult, "oneaudit:") {
				baseEndpoints = append(baseEndpoints, strings.Split(extractedResult, "oneaudit:")[1])
			} else {
				extractedCleaned = append(extractedCleaned, extractedResult)
			}
		}

		for _, extractedResult := range extractedCleaned {
			// We consider a duplicate when we found the same extract result for the same matcher and template
			// (like two URLs on a website having the exact same issue, e.g. a missing header)
			key := fmt.Sprintf("[%s:%s:%s]", result.TemplateID, result.MatcherName, extractedResult)

			// We don't modify the result as it is shared between extract results
			value := &ParsedEvent{
				Name:      extractedResult,
				Result:    result,
				Endpoints: append([]string{}, baseEndpoints...),
				Count:     0,
			}

			// Find the result or create it
			if _, ok := endpointsMap[key]; !ok {
				endpointsMap[key] = value
			} else {
				value = endpointsMap[key]
			}

			// Add one to the count (which is always >=) to the number of endpoints
			value.Count += 1

			// We need to compute a list of endpoints
			// (e.g. the list of URLs in which we found the same extract "key")
			newEndpoint := extractEndpoint(&result, value)
			if newEndpoint != "" {
				value.Endpoints = append(value.Endpoints, newEndpoint)
			}
		}
	}

	return endpointsMap
}

func extractEndpoint(result *output.ResultEvent, value *ParsedEvent) string {
	if result.Matched == "" {
		return ""
	}

	// Attempt to parse the URL
	parsedUrl, err := url.Parse(result.Matched)
	var newValue string
	if err == nil {
		newValue = parsedUrl.Path
	} else {
		newValue = result.Matched
	}
	// If there is no match
	if newValue == "" || !strings.HasPrefix(newValue, "/") {
		return ""
	}

	// Ensure the URL is not in the list
	for _, endpoint := range value.Endpoints {
		if endpoint == newValue {
			return ""
		}
	}

	return newValue
}
