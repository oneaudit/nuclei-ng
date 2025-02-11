package runner

import (
	"github.com/oneaudit/nuclei-ng/pkg/types"
	errorutil "github.com/projectdiscovery/utils/errors"
	"os/exec"
)

func validateOptions(_ *types.Options) error {
	cmd := exec.Command("nuclei", "-version")
	err := cmd.Run()
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("could not execute nuclei")
	}
	return nil
}
