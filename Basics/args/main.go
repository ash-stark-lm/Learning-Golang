package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echoClassic()
	echoRange()
	echoJoin()
	echoDebug()
}

// Version 1: classic for loop with index
func echoClassic() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Version 2: range-based loop
func echoRange() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Version 3: using strings.Join (most efficient)
func echoJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echoDebug() {
	fmt.Println(os.Args[1:])
}
