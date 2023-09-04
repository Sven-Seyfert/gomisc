package clipboard

import (
	"fmt"
	"os/exec"
	"strings"
)

func read() (string, error) {
	// Command looks for a file name, in case a file was copied to the clipboard.
	command := "Get-Clipboard -Format FileDropList -RAW"
	cmd := exec.Command("powershell", command)

	content, err := getClipboardText(cmd)
	if err != nil {
		return "", err
	}

	if content != "" {
		// Return found file path.
		return removeTrailingNewline(content), nil
	}

	// Command simply gets the clipboard text.
	command = "Get-Clipboard -Format text"
	cmd = exec.Command("powershell", command)

	content, err = getClipboardText(cmd)
	if err != nil {
		return "", err
	}

	// Return found text.
	return removeTrailingNewline(content), nil
}

func getClipboardText(cmd *exec.Cmd) (string, error) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	defer stdin.Close()

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func removeTrailingNewline(text string) string {
	return strings.TrimSuffix(text, "\r\n")
}

func write(text string) error {
	// Command uses a powershell here-string to write to the clipboard.
	command := fmt.Sprintf("$rawString = @'\n%s\n'@; Set-Clipboard -Value $rawString", text)
	cmd := exec.Command("powershell", command)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
