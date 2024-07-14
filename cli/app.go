package cli

import (
	"fmt"

	"github.com/kerogs/KerogsGo/cli"
	"github.com/kerogs/KerogsGo/colors"
)

// HelloShow prints the Minecraft-Server-AutoCreation version to the console in red color.
//
// Parameters:
// - version: the version of the Minecraft-Server-AutoCreation application.
//
// Return type: None.
func HelloShow(version string) {
	cli.AsciiStart()
	fmt.Println(colors.Red + "Minecraft-Server-AutoCreation v" + version + colors.Reset)
}