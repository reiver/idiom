package main

import (
	_ "./commands"
	"./kernel"
	"./string"

	"bufio"
	"fmt"
	"os"
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
		scanner := bufio.NewReader(file)


		var lineNumber int64 = 0
		var result idiom_string.Type
		for {
			lineNumber++

			tokens := readline(scanner)
			if 0 >= len(tokens) {
		/////////////// CONTINUE
				continue
			}

			token0 := tokens[0]
			if idiom_string.Error("end of file") == token0 {
				os.Exit(0)
			}

			var command string
			{
				var err error
				command, err = token0.Return()
				if nil != err {
panic(err)
					
				}
			}

			parameters := tokens[1:]

			result = idiom_kernel.Run(command, parameters...)
			if result.IsError() {
				fmt.Fprintln(os.Stderr, "\n"+"\x1b[30;41m"+"ERROR"+"\x1b[0m"+" on line", lineNumber, "for command", command, ":", result.Err())
				os.Exit(1)
				return
			}
		}
	}
}
