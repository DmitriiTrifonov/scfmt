package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// TODO: Add a call for folders

	cannotParse := func(args interface{}) {
		log.Fatalln("cannot parse:", args)
	}

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		cannotParse(args)
	}

	if len(args) > 1 {
		cannotParse(args)
	}

	file, err := os.Open(filepath.Clean(args[0]))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

}
