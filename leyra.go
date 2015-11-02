package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/leyra/color.v1"
)

// TODO: We'll probably not use the master branch for this eventually
const url = "https://github.com/leyra/leyra/archive/master.zip"

type leyra struct {
	url     string
	archive string
}

func (l *leyra) get(args []string) {
	l.url = url
	l.download()
	l.unzip()
	l.setupDirs(args[0])
	l.rename("leyra-master", args[0]+"/src/leyra")
}

func (l *leyra) fetch(args []string) {
	l.url = fmt.Sprintf("https://github.com/%s/archive/master.zip", args[0])
	parts := strings.Split(args[0], "/")
	name := parts[1]
	l.download()
	l.unzip()
	l.setupDirs(name)
	l.rename(name+"-master", name+"/src/leyra")
}

func (l *leyra) download() {
	// Gets the file tokens
	t := strings.Split(l.url, "/")
	l.archive = t[len(t)-1]

	output, err := os.Create(l.archive)
	if err != nil {
		fmt.Println("Error while creating", l.archive, "-", err)
		return
	}

	defer output.Close()

	response, err := http.Get(l.url)
	if err != nil {
		fmt.Println("Error while downloading", l.url, "-", err)
		return
	}

	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", l.url, "-", err)
		return
	}

	fmt.Println("Finished!")
}

func (l leyra) unzip() {
	reader, err := zip.OpenReader(l.archive)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer reader.Close()

	for _, f := range reader.Reader.File {
		zipped, err := f.Open()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer zipped.Close()

		path := filepath.Join("./", f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			defer writer.Close()

			if _, err = io.Copy(writer, zipped); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			green := color.New(color.FgGreen).SprintFunc()
			fmt.Printf("%s %s\n", green("create"), path)
		}
	}
	os.Remove("master.zip")
}

func (l leyra) rename(d, dest string) {
	os.Rename(d, dest)
}

func (l leyra) setupDirs(directoryName string) {
	os.Mkdir(directoryName, 0755)
	os.Mkdir(directoryName+"/src", 0755)
}
