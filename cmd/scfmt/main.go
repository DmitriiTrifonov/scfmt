package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type codeLine struct {
	lineType lineType
	line     string
}

func (c *codeLine) String() string {
	return fmt.Sprintf("%d: %s", c.lineType, c.line)
}

type lineType uint

const (
	lineType_Invalid lineType = iota
	lineType_Pre
	lineType_OpenBracket
	lineType_CloseBracket
	lineType_Comma
	lineType_Terminated
	lineType_Empty
)

func main() {
	// TODO: Add a call for folders

	cannotParse := func(args interface{}) {
		log.Fatalln("cannot parse:", args)
	}
	isStdOutPtr := flag.Bool("stdout", false, "write to stdout")
	flag.Parse()

	isStdOut := *isStdOutPtr

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

	programText := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		line := newCodeLine(text)
		switch line.lineType {
		case lineType_Invalid:
			programText = append(programText, text+";")
		default:
			programText = append(programText, text)
		}
	}

	file.Close()
	err = os.Remove(filepath.Clean(args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if !isStdOut {
		output, err := os.Create(filepath.Clean(args[0]))
		if err != nil {
			log.Fatal(err)
		}
		defer output.Close()
		for _, line := range programText {
			_, err := fmt.Fprintf(output, "%s\n", line)
			if err != nil {
				log.Fatal(err)
			}
		}
		return
	}

	for _, line := range programText {
		fmt.Println(line)
	}

}

func newCodeLine(s string) codeLine {
	withoutSpaces := strings.TrimSpace(s)
	lineType := lineType_Invalid
	switch {
	case strings.HasPrefix(withoutSpaces, "#"):
		lineType = lineType_Pre
	case strings.HasSuffix(withoutSpaces, "{"):
		lineType = lineType_OpenBracket
	case strings.HasSuffix(withoutSpaces, "}"):
		lineType = lineType_CloseBracket
	case strings.HasSuffix(withoutSpaces, ","):
		lineType = lineType_Comma
	case len(withoutSpaces) == 0:
		lineType = lineType_Empty
	case strings.HasSuffix(withoutSpaces, ";"):
		lineType = lineType_Terminated
	default:
		lineType = lineType_Invalid
	}
	return codeLine{line: s, lineType: lineType}
}
