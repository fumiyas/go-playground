package main

import (
	"log"
	"os"
	"bufio"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("Usage: %s FILE\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	rd := bufio.NewReaderSize(file, 10)
	for {
		line, prefix_p, err := rd.ReadLine()
		if err != nil {
			break
		}

		var eol string
		if prefix_p {
			eol = ""
		} else {
			eol = "\n"
		}

		fmt.Printf("%s%s", line, eol)
	}
}

