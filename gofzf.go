package gofzf

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Select takes a slice of strings and returns a single string
// chosen by the user using fzf. It's strongly recommended to trim spaces
// of strings in input.
func Select(input []string) (string, error) {
	var fzfInput strings.Builder
	for _, v := range input {
		fzfInput.WriteString(fmt.Sprintf("%s\n", v))
	}
	cmd := exec.Command("fzf")
	var result strings.Builder
	cmd.Stdout = &result
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(stdin, strings.NewReader(fzfInput.String()))
	if err != nil {
		return "", err
	}
	err = stdin.Close()
	if err != nil {
		return "", err
	}
	err = cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(result.String()), nil
}
