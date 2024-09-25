package ansible

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
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

func WriteConfigToFile(url string, tempDir string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data from URL: %s", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read data: %s", err)
	}

	filePath := filepath.Join(tempDir, "vars.yaml")
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write configuration file: %s", err)
	}
	fmt.Println("Configuration file successfully written:", filePath)
}
