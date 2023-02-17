package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("This is the contents written to file!")
	err := ioutil.WriteFile("file", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
