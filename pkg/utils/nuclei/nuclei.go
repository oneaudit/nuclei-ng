package nuclei

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	openapiutil "github.com/oneaudit/nuclei-ng/pkg/utils/openapi"
	"github.com/projectdiscovery/gologger"
	errorutil "github.com/projectdiscovery/utils/errors"
	"math"
	"os"
	"os/exec"
	"strconv"
)

func ExecuteCommand(options *types.Options, tags types.Tag, specification *openapi3.T, paths *openapi3.Paths) ([]byte, error) {
	tempFile, err := openapiutil.CreateTemporarySwaggerFile(specification, paths)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoDeferInLoop,GoUnhandledErrorResult
	defer os.Remove(tempFile.Name())

	// Run the command
	cmd := exec.Command("nuclei",
		"-dast", "-no-interactsh",

		"-im", "openapi",
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
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errorutil.NewWithErr(err).Msgf("%s", cmdOutput)
	}

	return cmdOutput, nil
}
