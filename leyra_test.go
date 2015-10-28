package main

import (
	"log"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	l := leyra{}
	l.get([]string{
		"test_app",
	})

	resApp, err := exists("test_app")
	if resApp != true {
		t.Error("Expecting to find 'test_app' directory. Does not exist.")
	}

	if err != nil {
		log.Fatal(err)
	}

	os.RemoveAll("test_app")

	resZip, err := exists("master.zip")
	if resZip != false {
		t.Error("master.zip found. File should be deleted")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
