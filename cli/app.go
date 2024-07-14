package cli

import (
	"fmt"

	"github.com/kerogs/KerogsGo/cli"
	"github.com/kerogs/KerogsGo/colors"
)

func HelloShow(version string) {
	cli.AsciiStart()
	fmt.Print(colors.Yellow)
	fmt.Println("#################################")
	fmt.Println("# Minecraft-Server-AutoCreation #")
	fmt.Println("# By Kerogs              v"+version+" #")
	fmt.Println("#################################")
	fmt.Print(colors.Reset+"\n")
}