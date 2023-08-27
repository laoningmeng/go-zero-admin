package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestGetLineAuthor(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	filenamem := "git.go"
	full := path + "/" + filenamem
	fmt.Println(full)
	author, err := GetLineAuthor(full, 16)
	fmt.Println(author, err)
}
