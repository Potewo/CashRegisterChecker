package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var (
	teststring   string = "hogefuga"
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

func TestGetHeaderTest(t *testing.T) {
	testFileName := "test.csv"
	exceptValue := []string{"a", "b", "c"}
	file, err := os.Create(testFileName)
	defer func() {
		if err := os.Remove(testFileName); err != nil {
			t.Log("failed to remove test file")
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Log("failed to create new file for test")
		t.Fatal(err)
	}
	defer file.Close()
	file.WriteString("a,b,c\n1,3,4")
	headerNames, err := getHeader(testFileName)
	t.Logf("header: %#v", headerNames)
	if !reflect.DeepEqual(headerNames, exceptValue) {
		t.Fatalf("not same except: %#v but: %#v", exceptValue, headerNames)
	}
}
