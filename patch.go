package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"golang.org/x/oauth2"
	"gopkg.in/leyra/go-github.v1/github"
)

type patch struct {
}

func (p *patch) authenticate(token string) {
	e, err := exists("./.token")
	if err != nil {
		log.Fatal(err)
	}

	if e != true {
		err := ioutil.WriteFile(".token", []byte(token), 0755)
		if err != nil {
			log.Fatal(err)
		}
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
	fmt.Printf("Patch generated - would you like to submit? (y / n) > ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	if text == "y\n" {
		e, err := exists(".token")
		if err != nil {
			log.Fatal(err)
		}

		if e != true {
			fmt.Println("You are not authenticated with GitHub yet.")
			os.Exit(1)
		}

		token, err := ioutil.ReadFile("./.token")
		if err != nil {
			log.Fatal(err)
		}

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: string(token)},
		)

		tc := oauth2.NewClient(oauth2.NoContext, ts)
		client := github.NewClient(tc)

		title := "request-pull: first request-pull test"
		body := string(out)
		assignee := "hazbo"

		input := &github.IssueRequest{
			Title:    &title,
			Body:     &body,
			Assignee: &assignee,
			Labels:   &[]string{},
		}

		issue, _, err := client.Issues.Create("leyra", "gresh", input)
		if err != nil {
			log.Fatal("Issues.Create returned error: %v", err)
		}

		if issue != nil {
			fmt.Println("Patch submitted, thanks!")
		}
	}
}
