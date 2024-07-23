package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ListAndTidyServicesModules() {
	servicesPath := "./services" // Adjust the path as necessary
	directories, err := os.ReadDir(servicesPath)
	if err != nil {
		fmt.Printf("Error reading services directory: %s\n", err)
		return
	}

	originalDir, err := os.Getwd() // Store the current working directory
	if err != nil {
		fmt.Printf("Error getting current directory: %s\n", err)
		return
	}

	for _, dir := range directories {
		if dir.IsDir() {
			modulePath := filepath.Join(servicesPath, dir.Name())
			err := os.Chdir(modulePath) // Change to the module directory
			if err != nil {
				fmt.Printf("Error changing to module directory %s: %s\n", modulePath, err)
				continue
			}

			fmt.Printf("Tidying module: %s\n", modulePath)
			cmd := exec.Command("go", "mod", "tidy")
			if output, err := cmd.CombinedOutput(); err != nil {
				fmt.Printf("Error tidying module %s: %s\n", modulePath, err)
				fmt.Println(string(output))
			} else {
				fmt.Printf("Successfully tidied module: %s\n", modulePath)
			}

			err = os.Chdir(originalDir) // Change back to the original directory
			if err != nil {
				fmt.Printf("Error changing back to original directory: %s\n", err)
				return
			}
		}
	}
}
