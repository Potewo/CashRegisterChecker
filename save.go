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
	// "sort"
	"strings"
	"time"
)

type SaveData struct {
	Date   time.Time `json:"date"`
	Caches []struct {
		Initialvalue int    `json:"initialValue"`
		Unitcost     int    `json:"unitCost"`
		Name         string `json:"name"`
		Value        int    `json:"value"`
	} `json:"caches"`
	Sales         int `json:"sales"`
	Otherservices []struct {
		Unitcost   int    `json:"unitCost"`
		N          int    `json:"n"`
		Ispositive bool   `json:"isPositive"`
		Name       string `json:"name"`
	} `json:"otherServices"`
	Unpaid []int `json:"unpaid"`
	Ins    []int `json:"ins"`
	Outs   []int `json:"outs"`
	Others []int `json:"others"`
}

// func save(jsonStr string) error {
// 	fileName, err := currentFile()
// 	if err != nil {
// 		return err
// 	}
// 	jsonMap, err := convertJsonToMap(jsonStr)
// 	if err != nil {
// 		return err
// 	}
// 	headers := make([]string, len(jsonMap))
// 	i := 0
// 	for key := range jsonMap {
// 		headers[i] = key
// 		i++
// 	}
// 	sort.Strings(headers)
// 	headerOk, err := checkHeader(fileName, headers)
// 	fileToWrite := fileName
// 	if headerOk {
// 	} else {
// 		// Create a new file
// 		// fileToWrite = ""
// 	}
// 	// Append to fileToWrite
// 	return nil
// }

func currentFile() (string, error) {
	bytes, err := ioutil.ReadFile("currentFile")
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(bytes), "\n"), nil
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

func convertJsonToStruct(jsonStr string) (SaveData, error) {
	var d SaveData
	if err := json.Unmarshal([]byte(jsonStr), &d); err != nil {
		return d, err
	}
	return d, nil
}
