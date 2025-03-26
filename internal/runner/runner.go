package runner

import (
	nucleiinternal "github.com/oneaudit/nuclei-ng/internal/nuclei"
	"github.com/oneaudit/nuclei-ng/pkg/api"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	"github.com/projectdiscovery/gologger"
	errorutil "github.com/projectdiscovery/utils/errors"
	"os"
	"sort"
)

func Execute(options *types.Options) error {
	options.ConfigureOutput()
	showBanner()

	if options.Version {
		gologger.Info().Msgf("Current version: %s", version)
		return nil
	}

	if err := api.ValidateOptions(options); err != nil {
		return errorutil.NewWithErr(err).Msgf("could not validate options")
	}

	endpointsMap, err := api.RunNucleiNG(options)
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("error while running nuclei-ng")
	}

	// write all requests
	writer := nucleiinternal.NewNGStandardWriter()

	// Sort keys
	keys := make([]string, 0, len(endpointsMap))
	for key := range endpointsMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Write the stdout
	for _, key := range keys {
		value := endpointsMap[key]

		// only add a teaser
		sort.Strings(value.Endpoints)
		if len(value.Endpoints) > 4 {
			value.Endpoints = value.Endpoints[:3]
			value.Endpoints = append(value.Endpoints, "...")
		}

		// Print to stdout
		_, _ = os.Stdout.Write(writer.FormatScreen(&value.Result, value.Endpoints))
		_, _ = os.Stdout.Write([]byte("\n"))
	}
	_, _ = os.Stdout.Write([]byte("\n"))

	return nil
}
