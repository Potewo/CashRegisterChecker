package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	teststring string = "hogefuga"
)

func TestCurrentFileSuccess(t *testing.T) {
	oldFile, err := ioutil.ReadFile("currentFile")
	if err != nil {
		t.Fatalf("failed to open file by ioutil %#v", err)
	}

	file, err := os.Open("currentFile")
	if err != nil {
		t.Fatalf("failed to open file by os%#v", err)
	}
	defer file.Close()
	defer file.Write(oldFile)
	file.WriteString(teststring)
	filename := currentFile()
	if filename != teststring {
		t.Logf("except: %v | but: %v", teststring, filename)
		t.Fatalf("failed to test: not same")
	}
}
