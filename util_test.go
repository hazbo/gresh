package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestConfigBufer(t *testing.T) {
	err := ioutil.WriteFile("test", []byte("Hello, World!"), 0755)
	if err != nil {
		log.Fatal(err)
	}
	contents := configBuffer("test")
	if string(contents) != "Hello, World!" {
		t.Error(
			"Expecting contents of 'test' to be Hello, World! Got",
			string(contents),
		)
	}

	os.Remove("test")
}
