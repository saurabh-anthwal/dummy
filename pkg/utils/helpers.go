package utils

import (
	"fmt"
	"os"
)

// MergeMaps merges two or more maps. It overwrites the keys in preceding maps in case of a conflict.
// Returns a new map with merged items.
func MergeMaps(ms ...map[string]string) map[string]string {
	res := make(map[string]string)
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// MakeDirs create full paths.
func MakeDirs(dirs ...string) error {
	for _, path := range dirs {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create dir %s, got error: %v", path, err)
		}
	}
	return nil
}

// StringIndexOf returns the index of item in given list, return -1 if not found.
func IndexOf(list []interface{}, item interface{}) int {
	for i, listItem := range list {
		if item == listItem {
			return i
		}
	}
	return -1
}
