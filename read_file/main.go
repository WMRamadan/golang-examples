package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	body, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(string(body))
}
