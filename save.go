package main

import (
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
