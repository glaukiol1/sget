package main

import (
	"log"
	"os"
)

func main() {
	url := parseArgs(os.Args)

	log.Println(request(url))
}
