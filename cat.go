package cat

import "strings"

func Meow() string {
	return "Meow"
}

func Purr() string {
	return "Purr"
}

func Scratch() string {
	return "Scratch"
}

func Nap(s string) string {
	return "Nap" + strings.ToUpper(s)
}
