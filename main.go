package main

import (
	"fmt"
	"log"
	"msac/cli"
	"time"

	"github.com/kerogs/KerogsGo/colors"
)

var (
	versionChoose string
)

const (
	AppVersion string = "1.0.2"
	spigotUrl  string = "https://getbukkit.org/download/spigot"
)

// main is the entry point of the program.
//
// It displays a welcome message, retrieves Spigot versions from the specified URL,
// prompts the user to choose a version, downloads the corresponding server jar,
// prepares the server for start, starts the server, accepts the EULA automatically,
// and finalizes the server installation and world creation.
//
// No parameters.
// No return values.
func main() {
	cli.HelloShow(AppVersion)

	// ? Get Spigot versions
	fmt.Println(colors.Green + "Connect to the following url : " + spigotUrl + colors.Reset)

	versions, err := cli.Spigot(spigotUrl)
	if err != nil {
		log.Fatal(err)
		fmt.Println(colors.Red + "ErrorDetected timePause activate (bypass Spigot protecction)")
		time.Sleep(5000)
	}

	fmt.Printf("\033[1A\033[K")

	versionCount := 0
	versionTotal := len(versions)
	for _, version := range versions {

		if versionCount == versionTotal-1 {
			fmt.Println(version)
		} else {
			fmt.Print(version + ", ")
		}
		versionCount++

	}

	// ? Ask version to download
	fmt.Print(colors.Orange + "Choose a version : " + colors.Reset)
	fmt.Scanln(&versionChoose)

	// ? Download server jar
	fmt.Println("Try download server jar -> https://download.getbukkit.org/spigot/spigot-" + colors.Magenta + versionChoose + colors.Reset + ".jar")
	cli.SpigotDownload("https://download.getbukkit.org/spigot/spigot-" + versionChoose + ".jar")
	fmt.Println(colors.Green + "Download completed successfully" + colors.Reset)

	// ? Prepare server
	cli.PrepareStart(versionChoose, AppVersion)
	fmt.Println("File start.bat create")

	// ? Start bat
	fmt.Println(colors.Orange + "WARNING : If you have an error, make sure you have the correct version of java installed on your machine!"+colors.Reset)
	fmt.Println("First launch...")

	cli.StartBat(versionChoose, AppVersion)

	fmt.Println(colors.Green + "First launch successfully completed" + colors.Reset)

	// ? Auto accept EULA
	fmt.Println("Eula automatically accepted")
	cli.AcceptEula()

	// ? Finalize
	fmt.Println("Finalize server installation and create world...")
	cli.StartBat(versionChoose, AppVersion)

}
