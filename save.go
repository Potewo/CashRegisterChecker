package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func save() {
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// output := "hello from go"
	// file.Write(([]byte)(output))
	// fileName := currentFile()
}

func currentFile() string {
	bytes, err := ioutil.ReadFile("currentFile")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(string(bytes), "\n")
}

func getHeader(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	record, err := reader.Read()
	if err == io.EOF {
		return make([]string, 1), err
	}
	if err != nil {
		log.Fatal(err)
	}
	return record, nil
}

func appendToFile(string) {
}
