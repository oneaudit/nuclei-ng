package nuclei

import (
	"bytes"
	"github.com/logrusorgru/aurora"
	"github.com/oneaudit/nuclei-ng/internal/colorizer"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/nuclei/v3/pkg/model/types/severity"
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
	"io"
	"strconv"
	"strings"
	"sync"

	"github.com/projectdiscovery/nuclei/v3/pkg/types"
	mapsutil "github.com/projectdiscovery/utils/maps"
)

type NGStandardWriter struct {
	json                  bool
	jsonReqResp           bool
	timestamp             bool
	noMetadata            bool
	matcherStatus         bool
	mutex                 *sync.Mutex
	aurora                aurora.Aurora
	outputFile            io.WriteCloser
	traceFile             io.WriteCloser
	errorFile             io.WriteCloser
	severityColors        func(severity.Severity) string
	storeResponse         bool
	storeResponseDir      string
	omitTemplate          bool
	DisableStdout         bool
	AddNewLinesOutputFile bool // by default this is only done for stdout
	KeysToRedact          []string
}

func NewNGStandardWriter() *NGStandardWriter {
	auroraColorizer := aurora.NewAurora(true)
	return &NGStandardWriter{
		json:             false,
		jsonReqResp:      true,
		noMetadata:       false,
		matcherStatus:    false,
		timestamp:        false,
		aurora:           auroraColorizer,
		mutex:            &sync.Mutex{},
		outputFile:       nil,
		traceFile:        nil,
		errorFile:        nil,
		severityColors:   colorizer.New(auroraColorizer),
		storeResponse:    false,
		storeResponseDir: "",
		omitTemplate:     true,
		KeysToRedact:     goflags.StringSlice{},
	}
}

// FormatScreen formats the output for showing on screen.
func (w *NGStandardWriter) FormatScreen(output *output.ResultEvent, endpoints []string) []byte {
	builder := &bytes.Buffer{}

	if !w.noMetadata {
		if w.timestamp {
			builder.WriteRune('[')
			builder.WriteString(w.aurora.Cyan(output.Timestamp.Format("2006-01-02 15:04:05")).String())
			builder.WriteString("] ")
		}
		builder.WriteRune('[')
		builder.WriteString(w.aurora.BrightGreen(output.TemplateID).String())

		if output.MatcherName != "" && !strings.HasPrefix(output.MatcherName, "_") {
			builder.WriteString(":")
			builder.WriteString(w.aurora.BrightGreen(output.MatcherName).Bold().String())
		} else if output.ExtractorName != "" && !strings.HasPrefix(output.MatcherName, "_") {
			builder.WriteString(":")
			builder.WriteString(w.aurora.BrightGreen(output.ExtractorName).Bold().String())
		}

		if w.matcherStatus {
			builder.WriteString("] [")
			if !output.MatcherStatus {
				builder.WriteString(w.aurora.Red("failed").String())
			} else {
				builder.WriteString(w.aurora.Green("matched").String())
			}
		}

		if output.GlobalMatchers {
			builder.WriteString("] [")
			builder.WriteString(w.aurora.BrightMagenta("global").String())
		}

		builder.WriteString("] [")
		builder.WriteString(w.aurora.BrightBlue(output.Type).String())
		builder.WriteString("] ")

		builder.WriteString("[")
		builder.WriteString(w.severityColors(output.Info.SeverityHolder.Severity))
		builder.WriteString("] ")
	}
	if output.Matched != "" {
		builder.WriteString(output.Matched)
	} else {
		builder.WriteString(output.Host)
	}

	// If any extractors, write the results
	if len(output.ExtractedResults) > 0 {
		builder.WriteString(" [")

		for i, item := range output.ExtractedResults {
			// trim trailing space
			// quote non-ascii and non printable characters and then
			// unquote quotes (`"`) for readability
			item = strings.TrimSpace(item)
			item = strconv.QuoteToASCII(item)
			item = strings.ReplaceAll(item, `\"`, `"`)

			builder.WriteString(w.aurora.BrightCyan(item).String())

			if i != len(output.ExtractedResults)-1 {
				builder.WriteRune(',')
			}
		}
		builder.WriteString("]")
	}

	if len(output.Lines) > 0 {
		builder.WriteString(" [LN: ")

		for i, line := range output.Lines {
			builder.WriteString(strconv.Itoa(line))

			if i != len(output.Lines)-1 {
				builder.WriteString(",")
			}
		}
		builder.WriteString("]")
	}

	// Write meta if any
	if len(output.Metadata) > 0 {
		builder.WriteString(" [")

		first := true
		// sort to get predictable output
		for _, name := range mapsutil.GetSortedKeys(output.Metadata) {
			value := output.Metadata[name]
			if !first {
				builder.WriteRune(',')
			}
			first = false

			builder.WriteString(w.aurora.BrightYellow(name).String())
			builder.WriteRune('=')
			builder.WriteString(w.aurora.BrightYellow(strconv.QuoteToASCII(types.ToString(value))).String())
		}
		builder.WriteString("]")
	}

	// If it is a fuzzing output, enrich with additional
	// metadata for the match.
	if output.IsFuzzingResult {
		if output.FuzzingParameter != "" {
			builder.WriteString(" [")
			builder.WriteString(output.FuzzingPosition)
			builder.WriteRune(':')
			builder.WriteString(w.aurora.BrightMagenta(output.FuzzingParameter).String())
			builder.WriteString("]")
		}

		builder.WriteString(" [")
		builder.WriteString(output.FuzzingMethod)
		builder.WriteString("]")
	}

	if len(endpoints) > 0 {
		builder.WriteString(" [")
		first := true
		for _, endpoint := range endpoints {
			if !first {
				builder.WriteString(", ")
			}
			first = false

			builder.WriteString(w.aurora.BrightWhite(endpoint).String())
		}
		builder.WriteString("]")
	}

	return builder.Bytes()
}
