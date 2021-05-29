package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	teststring string = "hogefuga"
	testfilename string = "currentFile"
)

func TestCurrentFileSuccess(t *testing.T) {
	oldFile, err := ioutil.ReadFile(testfilename)
	if err != nil {
		t.Fatalf("failed to open file by ioutil %#v", err)
	}

	file, err := os.Create(testfilename)
	if err != nil {
		t.Fatalf("failed to open file by os%#v", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			t.Logf("failed to close file")
		}
		file, err := os.Create(testfilename)
		if err != nil {
			t.Logf("failed to reopen file")
		}
		file.Write(oldFile)
		t.Logf("Writed %#v", string(oldFile))
		err = file.Close()
		if err != nil {
			t.Logf("failed to close file")
		}
	}()
	_, err = file.WriteString(teststring)
	if err != nil {
		t.Fatalf("failed to Writing file")
	}
	filename := currentFile()
	if filename != teststring {
		t.Logf("except: %#v | but: %#v", teststring, filename)
		t.Fatalf("failed to test: not same")
	}
}
