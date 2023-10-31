package execbash

import (
	"bytes"
	"os/exec"
)

func execBash(text string) (string, string, error) {
	cmd := exec.Command("bash", "-c", text)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
