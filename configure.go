package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"gopkg.in/leyra/toml.v1"
)

var configFiles = []string{
	"deps.conf",
	"rc.conf",
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

type makefileConf struct {
	Deps struct {
		Goget            []string
		EnableLocalStubs string
	}
}

type makefileStub struct {
	Goget            []string
	EnableLocalStubs string
}

// TODO: This whole function is a mess and parts of it will be needed for all
// other stubs, so although it works, it needs fixing. The naming conventions
// here are awful. And we shouldn't need two *almost* identicle structs that
// hold the same data.
func makefileFromStub() {
	sf, err := ioutil.ReadFile("./extras/stubs/Makefile")
	if err != nil {
		panic(err)
	}

	var makefileC makefileConf
	buf := configBuffer("./etc/deps.conf")
	if err := toml.Unmarshal(buf, &makefileC); err != nil {
		panic(err)
	}

	makefileS := makefileStub{
		Goget:            makefileC.Deps.Goget,
		EnableLocalStubs: makefileC.Deps.EnableLocalStubs,
	}

	t := template.Must(template.New("letter").Parse(string(sf)))
	err = t.Execute(os.Stdout, makefileS)
}
