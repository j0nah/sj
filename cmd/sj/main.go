package main

import (
	"fmt"
	"github.com/j0nah/sj"
	"io/ioutil"
	"log"
	"os"
)

const Usage = `
  Usage:
    sj <file> <key>
`

func main() {
	args := os.Args
	if len(args) != 3 {
		usage()
		return
	}

	haystack, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	needles, err := sj.Search(args[2], string(haystack))
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range *needles {
		fmt.Println(n)
	}
}

func usage() {
	fmt.Println(Usage)
}
