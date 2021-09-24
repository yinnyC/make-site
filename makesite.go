package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, world!")
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(fileContents))
}
