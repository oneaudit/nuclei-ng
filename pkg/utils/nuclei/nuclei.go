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
	var (
		tempFile    *os.File
		inputFormat string
		err         error
	)
	if tags != types.JsExt {
		inputFormat = "openapi"
		tempFile, err = openapiutil.CreateTemporarySwaggerFile(specification, paths)
	} else {
		inputFormat = "list"
		tempFile, err = openapiutil.CreateTemporaryEndpointsFile(paths)
	}

	if err != nil {
		return "", err
	}
	//goland:noinspection GoDeferInLoop,GoUnhandledErrorResult
	defer os.Remove(tempFile.Name())

	// Run the command
	cmd := exec.Command("nuclei",
		"-dast", "-no-interactsh",

		"-im", inputFormat,
		"-list", tempFile.Name(),

		"-disable-update-check",
		"-t", options.NucleiTemplateDir,
		"-update-template-dir", options.NucleiTemplateDir,

		"-tags", string(tags),

		"-follow-redirects", "-no-mhe",
		"-fuzz-param-frequency", strconv.Itoa(math.MaxInt32-1),

		"-j", "-silent", "-omit-raw", "-omit-template", "-no-color",

		"-config", options.NucleiConfig,
	)
	gologger.Debug().Msgf("Executing command: %v", cmd)
	return "", errorutil.New("Executing command: %v", cmd)

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

		for _, extractedResult := range extracted {
			// We consider a duplicate when we found the same extract result for the same matcher and template
			// (like two URLs on a website having the exact same issue, e.g. a missing header)
			key := fmt.Sprintf("[%s:%s:%s]", result.TemplateID, result.MatcherName, extractedResult)

			// We don't modify the result as it is shared between extract results
			value := &ParsedEvent{
				Name:      extractedResult,
				Result:    result,
				Endpoints: []string{},
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
	if newValue == "" {
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
