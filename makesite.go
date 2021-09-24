package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Text struct {
	Content string
}

func main() {
	var filename string

	flag.StringVar(&filename, "f", "", "name of file to write to html")
	flag.StringVar(&filename, "file", "", "name of file to write to html")

	flag.Parse()

	fileContent := readFile(filename)
	fileToWrite := strings.SplitN(filename, ".", 2)[0]

	writeTemplate("template.tmpl", fileToWrite, string(fileContent))
}

func readFile(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
func writeTemplate(tmpl, filename, fileContent string) {
	text := Text{fileContent}

	htmlFile, osErr := os.Create(filename + ".html")
	if osErr != nil {
		log.Fatal(osErr)
	}

	t := template.Must(template.ParseFiles(tmpl))
	execErr := t.Execute(htmlFile, text)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
