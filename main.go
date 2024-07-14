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
	AppVersion string = "1.0.0"
	spigotUrl  string = "https://getbukkit.org/download/spigot"
)

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
	fmt.Println("\n" + colors.Orange + "[Downloading server jar...]" + colors.Reset)
	fmt.Println("Try download server jar -> https://download.getbukkit.org/spigot/spigot-" + colors.Magenta + versionChoose + colors.Reset + ".jar")
	cli.SpigotDownload("https://download.getbukkit.org/spigot/spigot-" + versionChoose + ".jar")
	fmt.Println(colors.Green + "Download completed successfully" + colors.Reset)

	// ? Prepare server
	fmt.Println("\n" + colors.Orange + "[Preparing server...]" + colors.Reset)
	cli.PrepareStart(versionChoose, AppVersion)
	fmt.Println(colors.Green + ".bat file created successfully" + colors.Reset)

	// ? Start bat
	fmt.Println("First launch of .bat")
	fmt.Println("WARNING : If you have an error, make sure you have the correct version of java installed on your machine!")

	fmt.Println("\n" + colors.Orange + "[STARTING...]" + colors.Reset)

	cli.StartBat()

	fmt.Println(colors.Green + "First launch successfully completed" + colors.Reset)

	fmt.Println("\n\n" + colors.Blue + "Fin de la v" + AppVersion + "...")
	var enter string
	fmt.Scanln(&enter)

}
