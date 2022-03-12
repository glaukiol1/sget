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
