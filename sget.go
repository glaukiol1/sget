package main

import (
	"fmt"
	"log"
)

func main() {
	var url string
	fmt.Print("Enter URL: (start with http(s) ")
	fmt.Scanln(&url)
	if isURL(url) {
		log.Println(request(url))
	} else {
		log.Fatalln("Incorrect URL.")
	}
}
