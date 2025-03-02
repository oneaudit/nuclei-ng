package runner

import (
	"github.com/oneaudit/nuclei-ng/pkg/types"
	errorutil "github.com/projectdiscovery/utils/errors"
	"os"
	"os/exec"
)

func validateOptions(options *types.Options) error {
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
