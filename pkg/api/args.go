package api

import (
	"github.com/oneaudit/nuclei-ng/pkg/types"
	"github.com/projectdiscovery/goflags"
)

func MakeFlagSet(options *types.Options, cfgFile *string) *goflags.FlagSet {
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`nuclei-ng ...`)

	flagSet.CreateGroup("input", "Target",
		flagSet.StringVarP(&options.InputFile, "input", "i", "", "openapi input file"),
	)

	flagSet.CreateGroup("nuclei", "Nuclei",
		flagSet.StringVarP(&options.NucleiTemplateDir, "templates-dir", "t", "", "path to the nuclei templates directory"),
		flagSet.StringVar(&options.NucleiConfig, "nuclei", "", "path to the nuclei configuration file"),
		flagSet.StringVarP(&options.ProxyHost, "nuclei-proxy", "p", "", "nuclei-ng is using an internal proxy by design, which can forward requests to another proxy."),
	)

	flagSet.CreateGroup("config", "Configuration",
		flagSet.StringVar(cfgFile, "config", "", "path to the nuclei-ng configuration file"),
		flagSet.StringSliceVarP(&options.NucleiTags, "tags", "nt", nil, "allowed tags. All tags are allowed if empty.", goflags.CommaSeparatedStringSliceOptions),
		flagSet.StringSliceVarP(&options.NucleiWorkflows, "workflows", "nw", nil, "allowed workflows. All workflows are allowed if empty.", goflags.CommaSeparatedStringSliceOptions),
	)

	flagSet.CreateGroup("output", "Output",
		flagSet.BoolVar(&options.Silent, "silent", false, "display output only"),
		flagSet.BoolVarP(&options.Verbose, "verbose", "v", false, "display verbose output"),
		flagSet.BoolVar(&options.Debug, "debug", false, "display debug output"),
		flagSet.BoolVar(&options.Version, "version", false, "display project version"),
	)

	return flagSet
}
