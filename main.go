package main

import (
	_ "./commands"
	"./kernel"
	"./string"

	"bytes"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	// The idiom interpreter (currently) expects all the idiom source code to be in a single file.
	//
	// The ‘filename’ variable will contain the name of that single file of idiom source code.
	var filename string
	{
		if 2 != len(os.Args) {
			fmt.Fprintf(os.Stderr, "USAGE:\n\t%s <filename>\n", os.Args[0])
			os.Exit(1)
			return
		}

		arg1 := os.Args[1]

		filename = arg1
	}

	// The ‘file’ variable will contain a reference to the open file for the single file of idiom source code
	// whose name is contained in the ‘filename’ variable.
	var file *os.File
	{
		var err error

		file, err = os.Open(filename)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: could not open file %q: %s\n", filename, err)
			os.Exit(1)
			return
		}
		defer file.Close()
	}

	// The for-loop here reads in a single line of code, and then tries to interpret it.
	{
		scanner := bufio.NewScanner(file)
		scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error){

			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}
			if atEOF {
				return len(data), data, nil
			}

			if i := bytes.IndexAny(data, "|\u000A\u000B\u000C\u000D\u0085\u2028\u2029"); i >= 0 {
				r, size := utf8.DecodeRune(data[i:])
				switch r {
				case '|':
					return i, data[0:i], nil
				default:
					return i + size, data[0:i], nil
				}
			}

			return
		})


		var lineNumber int64 = 0
		var result idiom_string.Type
		for scanner.Scan() {
			lineNumber++

			line := scanner.Text()
			line = strings.Trim(line, " \t")

			// Skip over any "empty" lines.
			if "" == strings.Trim(line, " \t") {
				continue
			}

			tokens := strings.Split(line, " ") //@TODO: This needs to be replaced.
			if 0 >= len(tokens) {
				continue
			}

			command    := tokens[0]
			parameters := tokens[1:]

			var p []idiom_string.Type = make([]idiom_string.Type, len(parameters))
			for i:=0; i<len(p); i++ {
				p[i] = idiom_string.Something(parameters[i])
			}

			result = idiom_kernel.Run(command, p...)
			if result.IsError() {
				fmt.Fprintln(os.Stderr, "\n"+"\x1b[30;41m"+"ERROR"+"\x1b[0m"+" on line", lineNumber, ":", result.Err())
				fmt.Fprintf(os.Stderr, "\t%q\n", scanner.Text())
				os.Exit(1)
				return
			}
		}
		if err := scanner.Err(); nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: problem scanning %q: %s\n", filename, err)
			os.Exit(1)
			return
		}
	}
}
