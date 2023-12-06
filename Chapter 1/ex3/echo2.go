package main

import "strings"

func Echo2(args []string) string {
	return strings.Join(args[1:], " ")
}
