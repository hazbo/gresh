package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	//"golang.org/x/oauth2"
	//"gopkg.in/leyra/go-github.v1/github"
)

type patch struct {
}

func (p *patch) authenticate(token string) {
	err := ioutil.WriteFile(".token", []byte("something"), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func (p patch) generate(url, branch string) {
	args := []string{
		"request-pull",
		"master",
		url,
		branch,
	}

	out, err := exec.Command("git", args...).Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
