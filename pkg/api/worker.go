package api

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oneaudit/nuclei-ng/pkg/javascript"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	nucleiutil "github.com/oneaudit/nuclei-ng/pkg/utils/nuclei"
	openapiutil "github.com/oneaudit/nuclei-ng/pkg/utils/openapi"
	proxytutil "github.com/oneaudit/nuclei-ng/pkg/utils/proxy"
	"github.com/projectdiscovery/gologger"
	errorutil "github.com/projectdiscovery/utils/errors"
	"net/http"
)

func RunNucleiNG(options *types.Options) error {
	spec, err := openapi3.NewLoader().LoadFromFile(options.InputFile)
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("error loading Swagger file")
	}

	entries, err := openapiutil.CategorizeRoutesByTags(spec)
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("could not categorize routes by tags")
	}

	// Proxy requests
	proxy := proxytutil.CreateProxy(options)
	go func() {
		gologger.Info().Msgf("Starting proxy on :8081...")
		if err := http.ListenAndServe(":8081", proxy); err != nil {
			gologger.Info().Msgf("Error starting proxy server: %s", err)
		}
	}()

	for tags, paths := range entries {
		if paths.Len() == 0 {
			continue
		}

		if !options.IsTagAllowed(tags) {
			continue
		}

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

		// fixme: ...
		println(len(endpointsMap))
	}

	// Sort keys
	//keys := make([]string, 0, len(xxx))
	//for key := range xxx {
	//	keys = append(keys, key)
	//}
	//sort.Strings(keys)

	// Clean results
	//for _, key := range keys {
	//	value := xxx[key]
	//
	//	// Sort Endpoints
	//	sort.Strings(value.Endpoints)
	//
	//	// Display only one extracted value
	//	if value.Name != "" {
	//		value.Result.ExtractedResults = []string{value.Name}
	//	}
	//
	//	// Change message shown
	//	value.Result.Matched = ""
	//	value.Result.Host = fmt.Sprintf("Found on %d URLs", value.Count)
	//
	//	// Do not display fake fuzzing information
	//	if !strings.Contains(value.Result.Info.Tags.String(), "fuzzing") {
	//		value.Result.IsFuzzingResult = false
	//	}
	//}

	// fixme: ...
	return nil
}
