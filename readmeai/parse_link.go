package main

import "strings"

func parse_link(link string) bool {
	if strings.HasPrefix(link, "https://github.com/") {
		return true
	} else {
		return false
	}
}
