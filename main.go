package main

import (
	"fmt"
	"strings"
)

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

func main() {
	totalLength, up := lenAndUpper("nico")
	fmt.Println(totalLength, up)

	// repeatMe("nico", "lynn", "dal", "marl", "flynn")
}