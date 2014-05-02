package main

import (
	"log"
	"os"
	"syscall"
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
		if os.IsNotExist(err) {
			log.Print("ENOENT")
		}
		if err.(*os.PathError).Err == syscall.ENOENT {
			log.Print("ENOENT")
		}
		log.Printf("ERROR: %s", err)
		os.Exit(1)
	}
	defer file.Close()

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

