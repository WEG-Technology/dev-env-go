package ansible

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func ExecuteAnsiblePlaybook(tempDir string) {
	fmt.Println("Executing Ansible playbook...")
	cmd := exec.Command("ansible-playbook", "-i", "inventory.ini", "playbook.yaml")
	cmd.Dir = tempDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatalf("Starting Ansible playbook failed: %s", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatalf("Ansible playbook execution failed: %s", err)
	}
}

func WriteConfigToFile(config string, tempDir string) {
	filePath := fmt.Sprintf("%s/vars.yaml", tempDir)
	err := os.WriteFile(filePath, []byte(config), 0644)
	if err != nil {
		log.Fatalf("Error writing configuration file: %s", err)
	}
	fmt.Println("Configuration file successfully written:", filePath)
}
