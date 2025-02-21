package javascript

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oneaudit/nuclei-ng/pkg/types"
	nucleiutil "github.com/oneaudit/nuclei-ng/pkg/utils/nuclei"
	"github.com/projectdiscovery/nuclei/v3/pkg/model"
	"github.com/projectdiscovery/nuclei/v3/pkg/model/types/severity"
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
	"path"
	"regexp"
	"strings"
)

var (
	JsS                    = `(`
	JsP0                   = `((?:v|\/)(\d+\.\d+\.\d+))`
	JsP1                   = `|((?:\/|@)(\d+\.\d+\.\d+)\/)`
	JsP2                   = `|((?:\.)(\d+\.\d+\.\d+)(?:-))`
	JsP3                   = `|((?:\/|\.)(\d+\.\d+\.\d+))`
	JsE                    = `)`
	relativeEndpointsRegex = regexp.MustCompile(JsS + JsP0 + JsP1 + JsP2 + JsP3 + JsE)

	versionInNameRegex = regexp.MustCompile(`(\d+\.\d+\.\d+)`)

	templateID = "javascript-library"
)

func AnalyzeExternalScripts(_ *types.Options, paths *openapi3.Paths) (map[string]*nucleiutil.ParsedEvent, error) {
	endpointsMap := make(map[string]*nucleiutil.ParsedEvent)

	for libName, libItem := range paths.Map() {
		libFileName := path.Base(strings.Split(libName, "?")[0])
		relativeMatches := relativeEndpointsRegex.FindAllStringSubmatch(libName, -1)

		var extractedVersion string
		for _, match := range relativeMatches {
			if len(match) < 2 {
				continue
			}
			// Compute Result
			extractedVersion = match[3]
			if extractedVersion == "" {
				extractedVersion = match[5]
			}
		}

		if extractedVersion == "" {
			versionMatch := versionInNameRegex.FindStringSubmatch(libFileName)
			if len(versionMatch) > 1 {
				extractedVersion = versionMatch[0]
			}
		}

		if extractedVersion == "" {
			extractedVersion = "unknown"
		}

		// Add to the list
		extractedResult := fmt.Sprintf("%s==%s", libFileName, extractedVersion)
		key := fmt.Sprintf("[%s:%s:%s]", templateID, "", extractedResult)
		endpointsMap[key] = &nucleiutil.ParsedEvent{
			Name: extractedResult,
			Result: output.ResultEvent{
				Info: model.Info{
					SeverityHolder: severity.Holder{
						Severity: severity.Info,
					},
				},
				TemplateID: templateID,
				Type:       "code",
			},
			Endpoints: libItem.Extensions["x-endpoints"].([]string),
		}
		endpointsMap[key].Count = len(endpointsMap[key].Endpoints)
	}

	return endpointsMap, nil
}
