// Service to use to parse json files
// We can make it an interface that we can use to parse different files
// We will limit it at simply reading the content from a file without formatting the data inside
package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// A method to get the content of a file
func GetFileContent(filepath string) ([]byte, error) {
	// Get file data in string
	dat, err := os.ReadFile(filepath)
	defer fmt.Println(string(dat))
	if err != nil {
		return nil, fmt.Errorf("error while reading file: %s", err)
	}

	return dat, nil
}

// A function that data from a file and returns slice of map of string and interface
func ParseJsonFile(filepath string) ([]map[string]interface{}, error) {
	// Get file data in string
	dat, err := GetFileContent(filepath)
	if err != nil {
		return nil, err
	}
	var j []map[string]interface{}
	if err := json.Unmarshal(dat, &j); err != nil {
		return nil, fmt.Errorf("error while parsing json : %s", err)
	}

	return j, nil
}

func ParseTxtFile(filepath string) ([]string, error) {
	// Get file data in string
	dat, err := GetFileContent(filepath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(dat), "\n")

	return lines, nil
}
