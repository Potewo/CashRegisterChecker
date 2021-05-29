package main

import (
	"io/ioutil"
	"log"
	"strings"
	// "encoding/csv"
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

func getHeader() []string {
	keys := make([]string, 100)
	return keys
}

func appendToFile(string) {
}
