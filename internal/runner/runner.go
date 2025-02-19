package runner

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	nucleiutil "github.com/oneaudit/nuclei-ng/pkg/utils/nuclei"
	openapiutil "github.com/oneaudit/nuclei-ng/pkg/utils/openapi"
	"github.com/projectdiscovery/gologger"
	errorutil "github.com/projectdiscovery/utils/errors"
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

	for tags, paths := range entries {
		if paths.Len() == 0 {
			continue
		}

		gologger.Info().Msgf("Running nuclei with tags: %v against %d targets\n", tags, paths.Len())

		result, err := nucleiutil.ExecuteCommand(options, tags, spec, paths)
		if err != nil {
			return errorutil.NewWithErr(err).Msgf("Error executing nuclei command: %s", err.Error())
		}

	}

	return nil
}
