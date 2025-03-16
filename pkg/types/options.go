package types

import (
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	logutil "github.com/projectdiscovery/utils/log"
	"strings"
)

type Options struct {
	// InputFile openapi input file
	InputFile string
	// NucleiTemplateDir path with nuclei templates to use
	NucleiTemplateDir string
	// NucleiConfig path with nuclei configuration to use
	NucleiConfig string
	// NucleiTags allowed tags. All tags are allowed if empty.
	NucleiTags goflags.StringSlice
	// NucleiWorkflows allowed workflows. All workflows are allowed if empty.
	NucleiWorkflows goflags.StringSlice
	// ProxyHost proxy to use for requests
	ProxyHost string
	// Silent shows only output
	Silent bool
	// Verbose specifies showing verbose output
	Verbose bool
	// Debug
	Debug bool
	// DebugDirPath
	DebugDirPath string
	// Version enables showing of tool version
	Version bool
}

func (options *Options) IsTagAllowed(tag Tag) bool {
	var targets []string
	tagStr := string(tag)
	if tag.IsWorkflow() {
		// No filtering
		if len(options.NucleiWorkflows) == 0 {
			return true
		}
		// Must be allowed
		targets = options.NucleiWorkflows
	} else {
		// No filtering
		if len(options.NucleiTags) == 0 {
			return true
		}

		// Must be allowed
		targets = options.NucleiTags
	}

	tagStr = string(tag)
	for _, t := range targets {
		if strings.TrimSpace(t) == tagStr {
			return true
		}
	}

	return false
}

// ConfigureOutput configures the output logging levels to be displayed on the screen
func (options *Options) ConfigureOutput() {
	if options.Silent {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	} else if options.Verbose {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelWarning)
	} else if options.Debug {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	} else {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelInfo)
	}

	logutil.DisableDefaultLogger()
}
