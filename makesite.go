package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Text struct {
	Content string
}

func main() {
	fileContent := readFile("first-post.txt")
	fmt.Print(string(fileContent))
	writeTemplate("template.tmpl", "first-post", string(fileContent))
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
