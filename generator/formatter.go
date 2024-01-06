package generator

import (
	"bytes"
	"os/exec"
	"strings"
)

// FormatGoCode formats Go code using the "gofmt" command.
func formatGoCode(code string) (string, error) {
	cmd := exec.Command("gofmt")
	var out bytes.Buffer
	cmd.Stdin = strings.NewReader(code)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
