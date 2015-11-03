package main

import (
	"log"
	"testing"
	//"os/exec"
)

func TestPatch_Authenticate(t *testing.T) {
	patch := new(patch)

	patch.authenticate("token")

	e, err := exists(".token")

	if err != nil {
		log.Fatal(err)
	}

	if e != true {
		t.Fatal("Expecting to find ./.token file. Not found")
	}
}
