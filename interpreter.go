package main

import (
	"fmt"
)

var bytes = make(map[int]int)
var ptr = 0

const max int = 256

func Interprete(source string) {
	var loopIndex int = -1

	// loop over each character in the string
	for i := 0; i < len(source); i++ {
		v := source[i]

		switch string(v) {
		case "[":
			loopIndex = i
		case "]":
			if bytes[ptr] > 0 {
				i = loopIndex - 1
			} else {
				loopIndex = -1
			}
		case "+":
			bytes[ptr]++
		case "-":
			bytes[ptr]--
		case ".":
			fmt.Print(string(rune(bytes[ptr])))
			//fmt.Print(bytes[ptr], "_")
		case ",":
			fmt.Print("Input for position [", ptr, "] : ")
			var n int
			fmt.Scanf("%d", &n)
			bytes[ptr] = n

		case ">":
			ptr++
			if ptr > max-1 {
				ptr = 0
			}
		case "<":
			ptr--
			if ptr < 0 {
				ptr = max - 1
			}
		}
	}

	spacer := "-"
	for range fmt.Sprint("Bytes Map", bytes) {
		spacer += "-"
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(spacer)

	fmt.Println("End of Program")
	fmt.Println("Bytes Map", bytes)
	fmt.Println("Pointer", ptr)

	fmt.Println(spacer)
	fmt.Println()
}
