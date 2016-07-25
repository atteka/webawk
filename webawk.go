package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"github.com/atteka/libwebawk"
	"net/http"
	"os"
)

func getBodies(files []string) (io.Reader, error) {
	if len(files) == 0 {
		reader := bufio.NewReader(os.Stdin)

		return reader, nil
	}

	resp, err := http.Get(files[0])

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + files[0] + "\"")
		return nil, err
	}

	return resp.Body, nil

}

func ParseArgs() (string, []string) {
	offset := 0
	files := make([]string, 0, 10)
	program := ""

	flag.StringVar(&program, "f", "", "a string to program-file")
	flag.Parse()
	if program == "" {
		program = flag.Args()[0]
		offset = 1
	}
	for offset < len(flag.Args()) {
		files = append(files, flag.Args()[offset])
	}
	return program, files
}

func main() {
	program, files := ParseArgs()
	match, action, err := libwebawk.ParseWebawkProgram(program)
	if err != nil {
		return
	}
	//fmt.Println("match: " + match)
	//fmt.Println("action: " + action)

	bodies, err := getBodies(files)
	if err != nil {
		return
	}

	libwebawk.Run(bodies, match, action)
}
