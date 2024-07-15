package helper

import (
	"fmt"
	"os"

	"github.com/kerogs/KerogsGo/colors"
)

func StopProgram(valueReturn error, message string) {
	fmt.Println(valueReturn)
	
	if(message != "") {
		fmt.Println(colors.Red+message+colors.Reset)
	}

	fmt.Println("\n\n\n"+colors.Red+"Press a key to stop the program"+colors.Reset)
	fmt.Scanln()

	os.Exit(0)
}