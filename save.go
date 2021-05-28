package main

import (
	"io/ioutil"
	"log"
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
	
}


func currentFile() string {
	bytes, err := ioutil.ReadFile("currentFile")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func getHeader() []string {
	keys := make([]string, 100)
	return keys
}

func appendToFile(string) {
}
