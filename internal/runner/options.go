package runner

import (
	"github.com/oneaudit/nuclei-ng/pkg/types"
	errorutil "github.com/projectdiscovery/utils/errors"
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

	return nil
}
