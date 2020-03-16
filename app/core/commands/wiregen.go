package commands

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
)

func WireGen(c *cli.Context) {
	pathToContainer := "app/startup/di"

	// Check if container exists
	if _, err := os.Stat(pathToContainer); os.IsNotExist(err) {
		fmt.Printf("%s does not exist", pathToContainer)
		return
	}

	// Generate wire-gen
	cmd := exec.Command("sh",
		"-c",
		fmt.Sprintf("cd %s; wire; mv wire_gen.go ../../build/di/", pathToContainer),
	)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error generating wire_gen from %s\n", pathToContainer)
		return
	}

	fmt.Println("Success")
}
