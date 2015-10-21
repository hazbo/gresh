package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type file struct {
	url  string
	name string
}

func (f *file) download(url string) {
	// Gets the file tokens
	t := strings.Split(url, "/")
	fileName := t[len(t)-1]

	fmt.Println("Downloading latest version of Leyra...")
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}

	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println("Finished!")
}

func (f *file) unzip(path string, to string) {
	reader, err := zip.OpenReader(path)
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
	os.Rename("leyra-master", to)
	os.Remove("master.zip")
}
