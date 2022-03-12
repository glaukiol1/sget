package main

import "strings"

func isURL(url string) bool {
	starts := []string{"http://", "https://"}
	if len(url) > 10 {
		if strings.HasPrefix(url, starts[0]) || strings.HasPrefix(url, starts[1]) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func parseArgs(arguments []string) string {
	if len(arguments) == 1 {
		panic("unsufficient command line arguments. Usage:\n\t sget <url>")
	}
	args := arguments[1:]

	if isURL(args[0]) {
		return args[0]
	} else {
		panic("incorrect URL")
	}
}
