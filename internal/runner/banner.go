package runner

import (
	"fmt"
	"github.com/projectdiscovery/gologger"
)

var version = "v1.0.10"

var banner = fmt.Sprintf(`
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/ ng  %s
`, version)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tgithub.com/oneaudit\n\n")
}
