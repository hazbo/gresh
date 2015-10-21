package main

import (
	"io/ioutil"
	"os"
)

func configBuffer(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return buf
}
