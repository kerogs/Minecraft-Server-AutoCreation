package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func PrepareStart(versionSpigot string, versionApp string) {
	file, err := os.Create("start.bat")
	if err != nil {
		fmt.Println("file creation error :", err)
		return
	}
	defer file.Close()

	fileContent := "@echo off \n"
    fileContent += "title Minecraft Spigot version "+versionSpigot+" - github.com/kerogs \n"
    fileContent += "java -Xmx1G -jar spigot-1.21.jar \n"
    fileContent += "echo #################################\n"
    fileContent += "echo # Minecraft-Server-AutoCreation #\n"
    fileContent += "echo # By Kerogs              v"+versionApp+" #\n"
    fileContent += "echo #################################\n"
	fileContent += "PAUSE\n"

	_, err = file.WriteString(fileContent)
    if err != nil {
        fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
        return
    }

}

func StartBat() {
	batFilePath := filepath.Join(filepath.Dir(os.Args[0]), "start.bat")

	_, err := exec.LookPath(batFilePath)
    if err != nil {
        fmt.Println("Le fichier", batFilePath, "n'existe pas ou n'est pas accessible.")
        return
    }

	cmd := exec.Command("CMD.exe", "/C", batFilePath)
    err = cmd.Run()
    if err != nil {
        fmt.Println("Erreur lors de l'exécution de start.bat avec CMD.exe :", err)
        return
    }

}