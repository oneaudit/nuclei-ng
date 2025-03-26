package api

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oneaudit/nuclei-ng/pkg/javascript"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	nucleiutil "github.com/oneaudit/nuclei-ng/pkg/utils/nuclei"
	openapiutil "github.com/oneaudit/nuclei-ng/pkg/utils/openapi"
	proxytutil "github.com/oneaudit/nuclei-ng/pkg/utils/proxy"
	"github.com/projectdiscovery/gologger"
	errorutil "github.com/projectdiscovery/utils/errors"
	"net/http"
	"sort"
	"strings"
)

func RunNucleiNG(options *types.Options) (nucleiutil.EventMap, error) {
	spec, err := openapi3.NewLoader().LoadFromFile(options.InputFile)
	if err != nil {
		return nil, errorutil.NewWithErr(err).Msgf("error loading Swagger file")
	}

	entries, err := openapiutil.CategorizeRoutesByTags(spec)
	if err != nil {
		return nil, errorutil.NewWithErr(err).Msgf("could not categorize routes by tags")
	}

	// Proxy requests
	proxy := proxytutil.CreateProxy(options)
	go func() {
		gologger.Info().Msgf("Starting proxy on :8081...")
		if err := http.ListenAndServe(":8081", proxy); err != nil {
			gologger.Info().Msgf("Error starting proxy server: %s", err)
		}
	}()

	var endpointsMap = make(nucleiutil.EventMap)
	for tags, paths := range entries {
		if paths.Len() == 0 {
			continue
		}

		if !options.IsTagAllowed(tags) {
			continue
		}

		var endpointsMapForTag nucleiutil.EventMap
		if tags == types.JsExt {
			endpointsMapForTag, err = javascript.AnalyzeExternalScripts(options, paths)
			if err != nil {
				return nil, errorutil.NewWithErr(err).Msgf("error while analyzing external scripts")
			}
		} else {
			result, err := nucleiutil.ExecuteCommand(options, tags, spec, paths)
			if err != nil {
				return nil, errorutil.NewWithErr(err).Msgf("error executing nuclei command")
			}

			endpointsMapForTag = nucleiutil.ParseResult(result)
		}

		nucleiutil.MergeResults(endpointsMapForTag, endpointsMap)
	}

	// Clean results
	for _, value := range endpointsMap {
		// Sort Endpoints
		sort.Strings(value.Endpoints)

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
	}

	return endpointsMap, nil
}
