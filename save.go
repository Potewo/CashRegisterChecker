package main

import (
	"log"
	"os"
	// "encoding/csv"
)

func save() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output := "hello from go"
	file.Write(([]byte)(output))
}


func currentFile() string {
	return ""
}

func getHeader() []string {
	keys := make([]string, 100)
	return keys
}

func appendToFile(string) {
}
