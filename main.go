package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(string) string

func UpperCase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefixer string) TransformFunc {
	return func(s string) string {
		return prefixer + " " + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("Hello, World!", UpperCase))
	fmt.Println(transformString("Hello, World!", Prefixer("Prefixer")))
}
