package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

func save() {
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

func appendToFile(fileName string, body string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, body)
	return nil
}

func checkHeader(fileName string, newHeader []string) (bool, error) {
	header, err := getHeader(fileName)
	if err != nil {
		return false, err
	}
	if !reflect.DeepEqual(newHeader, header) {
		return false, nil
	}
	return true, nil
}

func convertJsonToMap(jsonStr string) (map[string]interface{}, error) {
	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &mapData); err != nil {
		return mapData, err
	}
	return mapData, nil
}
