package main

import (
	"fmt"
	"os"
)

var configFiles = []string{
	"database.conf",
	"deps.conf",
	"rc.conf",
	"server.conf",
}

func findConfigFiles() bool {
	for _, v := range configFiles {
		if findConfigFile(v) != true {
			fmt.Printf("Error: could not find: ./etc/%s\n", v)
			return false
		}
	}
	return true
}

func findConfigFile(fileName string) bool {
	if _, err := os.Stat("./etc/" + fileName); err == nil {
		return true
	}
	return false
}
