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

func TestFetch(t *testing.T) {
	l := leyra{}
	l.fetch([]string{
		"leyra/blog",
	})

	resApp, err := exists("blog")
	if resApp != true {
		t.Error("Expecting to find 'blog' directory. Does not exist.")
	}

	if err != nil {
		log.Fatal(err)
	}

	os.RemoveAll("blog")

	resZip, err := exists("master.zip")
	if resZip != false {
		t.Error("master.zip found. File should be deleted")
	}

	if err != nil {
		log.Fatal(err)
	}
}
