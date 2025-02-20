package runner

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	nucleiinternal "github.com/oneaudit/nuclei-ng/internal/nuclei"
	"github.com/oneaudit/nuclei-ng/pkg/javascript"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	nucleiutil "github.com/oneaudit/nuclei-ng/pkg/utils/nuclei"
	openapiutil "github.com/oneaudit/nuclei-ng/pkg/utils/openapi"
	"github.com/projectdiscovery/gologger"
	errorutil "github.com/projectdiscovery/utils/errors"
	"os"
	"sort"
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
		return errorutil.NewWithErr(err).Msgf("error loading Swagger file")
	}

	entries, err := openapiutil.CategorizeRoutesByTags(spec)
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("could not categorize routes by tags")
	}

	// write all requests
	writer := nucleiinternal.NewNGStandardWriter()

	for tags, paths := range entries {
		if paths.Len() == 0 {
			continue
		}

		gologger.Info().Msgf("Running nuclei with tags: [%v] against %d targets\n", tags, paths.Len())

		var endpointsMap map[string]*nucleiutil.ParsedEvent

		if tags == types.JsExt {
			endpointsMap, err = javascript.AnalyzeExternalScripts(options, paths)
			if err != nil {
				return errorutil.NewWithErr(err).Msgf("error while analyzing external scripts")
			}
		} else {
			result, err := nucleiutil.ExecuteCommand(options, tags, spec, paths)
			if err != nil {
				return errorutil.NewWithErr(err).Msgf("error executing nuclei command")
			}

			endpointsMap = nucleiutil.ParseResult(result)
		}

		// Sort keys
		keys := make([]string, 0, len(endpointsMap))
		for key := range endpointsMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		// Write the stdout
		for _, key := range keys {
			value := endpointsMap[key]

			// Add a teaser
			sort.Strings(value.Endpoints)
			if len(value.Endpoints) > 3 {
				value.Endpoints = value.Endpoints[:2]
				value.Endpoints = append(value.Endpoints, "...")
			}

			// Display only one extracted value
			if value.Name != "" {
				value.Result.ExtractedResults = []string{value.Name}
			}

			// Change message shown
			value.Result.Matched = ""
			value.Result.Host = fmt.Sprintf("Found on %d URLs", value.Count)

			// Do not display fake fuzzing information
			if !strings.Contains(value.Result.Info.Tags.String(), "fuzzing") {
				value.Result.IsFuzzingResult = false
			}

			// Print to stdout
			_, _ = os.Stdout.Write(writer.FormatScreen(&value.Result, value.Endpoints))
			_, _ = os.Stdout.Write([]byte("\n"))
		}
		_, _ = os.Stdout.Write([]byte("\n"))
	}

	return nil
}
