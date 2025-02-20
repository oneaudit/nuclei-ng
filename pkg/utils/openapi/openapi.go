package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	"github.com/oneaudit/nuclei-ng/pkg/utils/extensions"
	"github.com/projectdiscovery/gologger"
	errorutil "github.com/projectdiscovery/utils/errors"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

func CategorizeRoutesByTags(specification *openapi3.T) (map[types.Tag]*openapi3.Paths, error) {
	var entries = make(map[types.Tag]*openapi3.Paths)
	for _, tag := range types.AllTags {
		entries[tag] = &openapi3.Paths{}
	}

	fakeItem := &openapi3.PathItem{
		Get: openapi3.NewOperation(),
	}
	fakeItem.Get.Responses = openapi3.NewResponses()

	for path, item := range specification.Paths.Map() {
		// All URLs can use GENERIC templates
		entries[types.GENERIC].Set(path, item)
		ext := filepath.Ext(path)

		// Files matching the HTML filter can use HTML templates
		if extensions.IsHTMLFile(ext) {
			entries[types.HTML].Set(path, item)
		}

		if strings.HasSuffix(ext, ".js") {
			if extensions.IsPathCommonJSLibraryFile(path) {
				// These files must not have been modified, right?
			} else {
				// These files must have secrets, right?
			}

			// Try to determine the version of JavaScript Libs
			entries[types.JsExt].Set(specification.Servers[0].URL+path, fakeItem)
		}

		// Try to determine the version of JavaScript Libs
		if libs, found := item.Extensions["x-javascript-libs"].([]interface{}); found {
			for _, _libName := range libs {
				libName := _libName.(string)
				entries[types.JsExt].Set(libName, fakeItem)
			}
		}
	}

	return entries, nil
}

func CreateTemporarySwaggerFile(specification *openapi3.T, paths *openapi3.Paths) (*os.File, error) {
	tempFile, err := os.CreateTemp("", "swagger.yaml")
	if err != nil {
		return nil, errorutil.NewWithErr(err).Msgf("error creating temp file")
	}
	//goland:noinspection GoUnhandledErrorResult
	defer tempFile.Close()

	gologger.Info().Msgf("Temporary file created: %s\n", tempFile.Name())
	encoder := yaml.NewEncoder(tempFile)
	encoder.SetIndent(2)
	specification.Paths = paths
	err = encoder.Encode(&specification)
	if err != nil {
		return nil, errorutil.NewWithErr(err).Msgf("could not write output file: %s", tempFile.Name())
	}
	return tempFile, nil
}
