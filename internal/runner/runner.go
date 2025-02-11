package runner

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	"github.com/projectdiscovery/gologger"
	errorutil "github.com/projectdiscovery/utils/errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/exec"
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
		return errorutil.NewWithErr(err).Msgf("Error loading Swagger file")
	}
	var entries = make(map[string]*openapi3.Paths)
	entries["generic"] = &openapi3.Paths{}
	entries["images"] = &openapi3.Paths{}

	for path, item := range spec.Paths.Map() {
		entries["generic"].Set(path, item)
		//if strings.HasSuffix(path, ".xml") {
		//entries["images"].Set(path, item)
		//}
	}

	for category, paths := range entries {
		if paths.Len() == 0 {
			continue
		}

		// Compute nuclei options
		var tags []string
		switch category {
		case "image":
			tags = append(tags, "image")
		case "generic":
			tags = append(tags, "generic")
		}

		// Create a temporary swagger file
		// With the paths associated to the tag
		tempFile, err := os.CreateTemp("", "swagger.yaml")
		if err != nil {
			return errorutil.NewWithErr(err).Msgf("Error creating temp file")
		}
		defer tempFile.Close()
		gologger.Info().Msgf("Temporary file created: %s\n", tempFile.Name())
		encoder := yaml.NewEncoder(tempFile)
		encoder.SetIndent(2)
		spec.Paths = paths
		err = encoder.Encode(&spec)
		if err != nil {
			return errorutil.NewWithErr(err).Msgf("could not write output file: %s", tempFile.Name())
		}
		defer os.Remove(tempFile.Name())

		// Run nuclei
		cmd := exec.Command("nuclei",
			"-dast", "-no-interactsh",

			"-im", "openapi",
			"-list", tempFile.Name(),

			"-disable-update-check",
			"-t", "../nuclei-dast-templates/",
			"-update-template-dir", "../nuclei-dast-templates/",

			"-tags", strings.Join(tags, " "),
		)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Error executing command: %v", err)
		}
		fmt.Println(string(output))
	}

	return nil
}
