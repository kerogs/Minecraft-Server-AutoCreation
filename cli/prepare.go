package cli

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kerogs/KerogsGo/colors"
)

// PrepareStart prepares the start.bat file for Minecraft server with the given Spigot version and application version.
//
// Parameters:
// - versionSpigot: the version of Spigot for the Minecraft server.
// - versionApp: the version of the Minecraft-Server-AutoCreation application.
func PrepareStart(versionSpigot string, versionApp string) {
	file, err := os.Create("start.bat")
	if err != nil {
		fmt.Println("file creation error :", err)
		return
	}
	defer file.Close()

	fileContent := "@echo off \n"
	fileContent += "title Minecraft Spigot version " + versionSpigot + " - github.com/kerogs \n"
	fileContent += "java -Xmx1G -jar spigot-" + versionSpigot + ".jar nogui\n"
	fileContent += "echo #################################\n"
	fileContent += "echo # Minecraft-Server-AutoCreation #\n"
	fileContent += "echo # By Kerogs              v" + versionApp + " #\n"
	fileContent += "echo #################################\n"
	fileContent += "PAUSE\n"

	fmt.Println("Configuration : java -Xmx1G -jar spigot-" + versionSpigot + ".jar")

	_, err = file.WriteString(fileContent)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}

}

// StartBat starts the execution of the start.bat file with the given version and versionApp.
//
// Parameters:
// - version: the version of the Minecraft server to be created.
// - versionApp: the version of the Minecraft-Server-AutoCreation application.
//
// No return values.
func StartBat(version string, versionApp string) {
	batFilePath := filepath.Join(filepath.Dir(os.Args[0]), "start.bat")

	_, err := exec.LookPath(batFilePath)
	if err != nil {
		fmt.Println("Le fichier", batFilePath, "n'existe pas ou n'est pas accessible.")
		return
	}

	cmd := exec.Command("CMD.exe", "/C", batFilePath)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Erreur lors de la création du pipe pour stdout :", err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Erreur lors de la création du pipe pour stderr :", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Erreur lors du démarrage de la commande :", err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "Server will start in 20 seconds") {
				fmt.Println("Server start in 20 seconds")
			} else if strings.Contains(line, "Starting minecraft server version "+version) {
				fmt.Println("Server creation in version " + version)
			} else if strings.Contains(line, "World Settings For [world]") {
				fmt.Println("Creation of the world.")
			} else if strings.Contains(line, "World Settings For [world_nether]") {
				fmt.Println("Creating the nether")
			} else if strings.Contains(line, "World Settings For [world_the_end]") {
				fmt.Println("Creating the end")
			} else if strings.Contains(line, "Preparing start region for dimension minecraft:overworld") {
				fmt.Println("Preparing the starting region for overworld")
			} else if strings.Contains(line, "Preparing start region for dimension minecraft:the_nether") {
				fmt.Println("Preparing the starting region for nether")
			} else if strings.Contains(line, "Preparing start region for dimension minecraft:the_end") {
				fmt.Println("Preparing the starting region for end")
			} else if strings.Contains(line, "For help, type \"help\"") {
				// ? Ready To Use
				fmt.Println(colors.Green + "Successfully created server, successfully generated world" + colors.Reset)

				fmt.Println()
				fmt.Println(" ##########################")
				fmt.Println(" # Server is ready to use #")
				fmt.Println(" # IP Local : localhost   #")
				fmt.Println(" # Port Local : 25565     #")
				fmt.Println(" ##########################")

                fmt.Println()
				fmt.Println("Minecraft-Server-AutoCreation v" + versionApp + " - github.com/kerogs")
                fmt.Println("You can close and delete the installation exe. and launch your server with start.bat")
			}
		}
	}()

	// TODO
	errScanner := bufio.NewScanner(stderr)
	go func() {
		for errScanner.Scan() {
			line := errScanner.Text()
			if strings.Contains(line, "Specific error message 1") {
				fmt.Fprintln(os.Stderr, "1")
			} else if strings.Contains(line, "Specific error message 2") {
				fmt.Fprintln(os.Stderr, "2")
			}
		}
	}()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Erreur lors de l'exécution de start.bat avec CMD.exe :", err)
	}
}

// AcceptEula updates the eula.txt file to set 'eula=true' if 'eula=false' is found.
//
// No parameters.
// No return values.
func AcceptEula() {
	input, err := ioutil.ReadFile("eula.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "eula=false") {
			lines[i] = "eula=true"
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile("eula.txt", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
