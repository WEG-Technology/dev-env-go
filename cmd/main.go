package main

import (
	"development-env/internal/ansible"
	"flag"
	"fmt"
	"os"
)

func main() {
	url := flag.String("url", "", "YAML file URL (required)")
	tempDir := flag.String("tempDir", "/tmp/development-environment", "Project directory")
	repoURL := flag.String("repoURL", "https://github.com/WEG-Technology/dev-env-ansible.git", "Repository URL")
	flag.Parse()

	if *url == "" {
		fmt.Println("YAML file URL is required")
		os.Exit(1)
	}

	ansible.InstallAnsible()
	ansible.PrepareWorkingDirectory(*tempDir)
	ansible.CloneRepository(*repoURL, *tempDir)

	ansible.WriteConfigToFile(*url, *tempDir)

	ansible.ExecuteAnsiblePlaybook(*tempDir)
}
