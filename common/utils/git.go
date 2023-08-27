package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetLineAuthor(filePath string, lineNumber int) (string, error) {
	cmd := exec.Command("git", "blame", "-L", fmt.Sprintf("%d,%d", lineNumber, lineNumber), "--porcelain", filePath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	authorLine := strings.TrimSpace(strings.Split(string(output), "\n")[0])
	a := string(output)
	fmt.Println(a)
	parts := strings.SplitN(authorLine, " ", 4)
	if len(parts) != 2 {
		return "", fmt.Errorf("无法解析作者行：%s", authorLine)
	}

	author := strings.TrimSpace(parts[1])
	return author, nil
}
