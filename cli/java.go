package cli

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

func Java() (string, error) {
	fmt.Println("Java verification...")
	// Command to check the version of Java
	cmd := exec.Command("java", "-version")

	// Capture the output
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Use a regular expression to find the version string
	re := regexp.MustCompile(`version "(\d+\.\d+\.\d+(_\d+)?)"`)
	matches := re.FindStringSubmatch(stderr.String())
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", errors.New("could not determine Java version")
}
