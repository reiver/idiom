package main

import (
	_ "./commands"
	"./kernel"
	"./string"

	"bufio"
	"fmt"
	"os"
	"strings"
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
		var lineNumber int64 = 0
		for scanner.Scan() {
			lineNumber++

			line := scanner.Text()

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

			result := idiom_kernel.Run(command, p...)
			if result.IsError() {
				fmt.Fprintln(os.Stderr, "\nERROR:", result.Err())
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
