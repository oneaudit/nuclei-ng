package types

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	logutil "github.com/projectdiscovery/utils/log"
)

type Options struct {
	// Silent shows only output
	Silent bool
	// Verbose specifies showing verbose output
	Verbose bool
	// Debug
	Debug bool
	// Version enables showing of tool version
	Version bool
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
