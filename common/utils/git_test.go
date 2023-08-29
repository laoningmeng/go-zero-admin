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
	filenamem := "logger.go"
	full := path + "/../logger/" + filenamem
	fmt.Println(full)
	author, _ := GetLineAuthor(full, 16)
	fmt.Println("auther:", author)
}

func TestOjbk(t *testing.T) {
	//today := time.Now().Format("2006-01-02")
	fmt.Printf("\u001B[31m remove :%d \u001B[34;1m add:%d \u001B[0m", 10, 20)
}
