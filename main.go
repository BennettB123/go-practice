package main

import (
	"fmt"
	"os"
)

const Padding rune = '='

func main() {
	if len(os.Args) < 2 {
		printUsageAndExit()
	}

	if os.Args[1] == "-d" || os.Args[1] == "--decode" {
		if len(os.Args) < 3 {
			printUsageAndExit()
		}
		_ = decode(os.Args[2])
	} else {
		output := encode(os.Args[1])
		fmt.Println(output)
	}
}

func printUsageAndExit() {
	fmt.Println("Usage:")
	fmt.Println("  Encoding:  ./base64 <text_to_encode>")
	fmt.Println("  Decoding:  ./base64 -d <text_to_decode>")
	os.Exit(1)
}
