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
