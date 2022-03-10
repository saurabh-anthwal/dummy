package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadJsonFile reads a json file at jsonPath and returns its content in a map.
func ReadJsonFile(jsonPath string) (map[string]interface{}, error) {
	var result map[string]interface{}

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed opening tfvars file: %s", err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("failed unmarshalling json: %s", err)
	}

	return result, nil
}
