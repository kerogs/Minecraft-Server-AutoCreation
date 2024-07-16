package main

import (
	"fmt"
	"log"
	"time"
	"errors"

	"github.com/kerogs/KerogsGo/colors"

	"github.com/kerogs/Minecraft-Server-AutoCreation/cli"
	"github.com/kerogs/Minecraft-Server-AutoCreation/helper"
)

var (
	versionChoose string
)

const (
	AppVersion string = "1.2.1"
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

	// ? Check if JAVA installed
	version, errs := cli.Java()
	if errs != nil {

		fmt.Print(colors.Orange+"Do you want to install Java ? (Y/n) : "+colors.Reset)
		dljava := ""
		fmt.Scanln(dljava)
		
			// ? DL JAVA or not
		if(dljava == "Y" || dljava == "y" || dljava == "yes" || dljava == "YES" || dljava == "Yes") {
		} else{
			helper.StopProgram(errors.New("java not installed, please install it -> https://www.oracle.com/fr/java/technologies/downloads/#jdk22-windows"), "You must install Java.")
		}
	} else {
		fmt.Println("Java version:", version + "\n Java installed, the server can be created")
	}

	// ? Get Spigot versions
	fmt.Println(colors.Green + "Connect to the following url : " + spigotUrl + colors.Reset)

	var versions []string // Déclarez versions comme une slice de chaînes
	var err error

	for {
		// Assumez que cli.Spigot(spigotUrl) retourne []string
		versions, err = cli.Spigot(spigotUrl)
		if err != nil {
			log.Println(err) // Utilisez log.Println pour ne pas arrêter le programme
			fmt.Println(colors.Red + "Error connecting to server. Retry in 3.5 seconds" + colors.Reset)
			time.Sleep(3500 * time.Millisecond)
		} else {
			break
		}
	}

	// Maintenant, versions est []string, vous pouvez l'utiliser comme tel
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
	fmt.Println(colors.Orange + "WARNING : If you have an error, make sure you have the correct version of java installed on your machine!" + colors.Reset)
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
