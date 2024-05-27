package ansible

import (
	"development-env/internal/util"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func InstallAnsible() {
	fmt.Println("Installing Ansible...")
	cmd := exec.Command("brew", "install", "ansible")
	out, errOut, err := util.RunCommand(cmd)

	if err != nil {
		if strings.Contains(errOut, "is already installed and up-to-date") {
			fmt.Println("Ansible is already installed and up-to-date.")
		} else {
			log.Fatalf("Ansible installation failed: %s\n%s", err, errOut)
		}
	} else {
		fmt.Println("Output:", out)
	}
}
