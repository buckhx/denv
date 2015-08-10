package api

import (
	"strings"
)

func Upper(arg string) string {
	return strings.ToUpper(arg)
}

func Lower(arg string) string {
	return strings.ToLower(arg)
}

func Echo(arg string) string {
	return arg
}
