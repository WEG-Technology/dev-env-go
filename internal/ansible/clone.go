package ansible

import (
	"development-env/internal/util"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func PrepareWorkingDirectory(tempDir string) {
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		if err := os.MkdirAll(tempDir, 0755); err != nil {
			log.Fatalf("Failed to create directory: %s", err)
		}
	}
	fmt.Println("Working directory prepared")
}

func CloneRepository(repoURL, tempDir string) {
	fmt.Println("Cloning Ansible repository...")

	if _, err := os.Stat(tempDir); !os.IsNotExist(err) {
		err := os.RemoveAll(tempDir)
		if err != nil {
			log.Fatalf("Failed to remove existing directory: %s\n", err)
		}
	}

	cmd := exec.Command("git", "clone", repoURL, tempDir)
	out, errOut, err := util.RunCommand(cmd)
	if err != nil {
		log.Fatalf("Cloning repository failed: %s\n%s", err, errOut)
	}
	fmt.Println("Output:", out)
}
