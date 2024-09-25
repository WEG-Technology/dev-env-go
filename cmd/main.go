package main

import (
	"development-env/internal/ansible"
	"development-env/internal/config"
	"flag"
	"fmt"
	"log"
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

	cfg, err := config.GetConfig(*url)
	if err != nil {
		log.Fatalf("Failed to get config: %s", err)
	}

	fmt.Println(config.ExportYaml(cfg))
	ansible.WriteConfigToFile(config.ExportYaml(cfg), *tempDir)

	ansible.ExecuteAnsiblePlaybook(*tempDir)
}
