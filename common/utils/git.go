package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

func GetLineAuthor(filePath string, lineNumber int) (string, error) {
	cmd := exec.Command("git", "blame", "-L", fmt.Sprintf("%d,%d", lineNumber, lineNumber), filePath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	info := string(output)
	reg, err := regexp.Compile(`\ \(([0-9a-zA-Z]+)\s`)
	res := reg.FindStringSubmatch(info)
	if len(res) > 0 {
		return res[1], nil
	}
	return "", errors.New("未找到git信息")
}
