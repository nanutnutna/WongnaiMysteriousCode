package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func reverse(s string) string {
	var word string
	for i := len(s) - 1; i >= 0; i-- {
		word += string(s[i])
	}
	return word
}

func main() {
	var whatIsIt, Sentence string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt = string(sd)
	for _, val := range strings.Split(whatIsIt, ":") {
		Sentence += val + " "
	}
	fmt.Println(strings.TrimSpace(reverse(Sentence)))
}
