package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/leyra/toml.v1"
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

type makefileStub struct {
	Deps struct {
		Goget            []string
		EnableLocalStubs string
	}
}

func makefileFromStub() {
	sf, err := ioutil.ReadFile("./extras/stubs/Makefile")
	if err != nil {
		panic(err)
	}

	var makefile makefileStub
	buf := configBuffer("./etc/deps.conf")
	if err := toml.Unmarshal(buf, &makefile); err != nil {
		panic(err)
	}

	// Needs removing, just for using var
	if string(sf) == "" {}
}
