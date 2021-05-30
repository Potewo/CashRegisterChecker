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

func TestAppendToFile(t *testing.T) {
	testFileName := "test.csv"
	expectedValue := "abc\ndef\n"
	oldString := "abc\n"
	file, err := os.Create(testFileName)
	if err != nil {
		t.Log("failed to create new file for test")
		t.Fatal(err)
	}
	defer func() {
		file.Close()
		if err := os.Remove(testFileName); err != nil {
			t.Log("failed to remove test file")
			t.Fatal(err)
		}
	}()
	file.WriteString(oldString)
	err = appendToFile(testFileName, "def")
	if err != nil {
		t.Logf("failed to run appendToFile()")
		t.Fatal(err)
	}
	if err = file.Close(); err != nil {
		t.Log("failed to close file")
		t.Fatal(err)
	}
	file, err = os.Open(testFileName)
	if err != nil {
		t.Log("failed to close file")
		t.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(file)
	t.Logf("output: %#v", string(bytes))
	if string(bytes) != expectedValue {
		t.Fatal("not same")
	}
}

func TestCheckHeader(t *testing.T) {
	expectedValue := []string{"a", "b", "c"}
	file, err := os.Create("test.txt")
	defer file.Close()
	if err != nil {
		t.Fatal("failed to create a new file")
	}
	file.WriteString("a,b,c\n1,2,3")
	file.Close()
	success, err := checkHeader("test.txt", expectedValue)
	if !success {
		t.Fatal("not same")
	}
	err = os.Remove("test.txt")
	if err != nil {
		t.Fatal(err)
	}
}
