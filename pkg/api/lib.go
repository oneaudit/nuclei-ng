package api

import (
	"github.com/oneaudit/nuclei-ng/pkg/types"
	errorutil "github.com/projectdiscovery/utils/errors"
	"os"
	"os/exec"
)

func CreateDefaultOptionsFromFile(cfgFile string) (*types.Options, error) {
	cfgFileOpts := ""
	cfgOptions := &types.Options{}
	flagSet := MakeFlagSet(cfgOptions, &cfgFileOpts)

	// Takes precedence over the parameter
	if cfgFile != "" {
		cfgFileOpts = cfgFile
	}

	if cfgFileOpts != "" {
		if err := flagSet.MergeConfigFile(cfgFileOpts); err != nil {
			return nil, errorutil.NewWithErr(err).Msgf("could not read config file")
		}
	}
	return cfgOptions, nil
}

func ValidateOptions(options *types.Options) error {
	cmd := exec.Command("nuclei", "-version")
	err := cmd.Run()
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("could not find nuclei")
	}

	if options.NucleiTemplateDir == "" {
		return errorutil.New("no nuclei template dir specified")
	}

	if options.Debug {
		options.DebugDirPath = ".debug"
		if _, err := os.Stat(options.DebugDirPath); os.IsNotExist(err) {
			err := os.Mkdir(options.DebugDirPath, os.ModePerm)
			if err != nil {
				return errorutil.New("Error creating .debug directory: %s", err.Error())
			}
		}
	}
	return nil
}
