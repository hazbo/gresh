package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/leyra/cli.v1"
)

const url = "https://github.com/leyra/leyra/archive/master.zip"

type leyra struct {
	url     string
	archive string
}

func (l leyra) get(c *cli.Context) {
	l.download()
	l.unzip()
	l.setupDirs(c.Args()[0])
	l.rename(c.Args()[0] + "/src/leyra")
}

func (l *leyra) download() {
	l.url = url

	// Gets the file tokens
	t := strings.Split(l.url, "/")
	l.archive = t[len(t)-1]

	fmt.Println("Downloading latest version of Leyra...")
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

			fmt.Println(path)

		}
	}
	os.Remove("master.zip")
}

func (l leyra) rename(dest string) {
	os.Rename("leyra-master", dest)
}

func (l leyra) setupDirs(directoryName string) {
	os.Mkdir(directoryName, 0755)
	os.Mkdir(directoryName+"/src", 0755)
}
