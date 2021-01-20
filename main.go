package main

import (
	_ "./commands"

	"./cmdln"
	"./kernel"
	"./log"
	"./parser"
	"./string"
	"./token"

	"fmt"
	"os"
)

func main() {

	idiom_log.Inform("[main] BEGIN ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")

	// The idiom interpreter (currently) expects all the idiom source code to be in a single file.
	//
	// The ‘filename’ variable will contain the name of that single file of idiom source code.
	var filename string
	{
		if 1 != len(idiom_cmdln.Args) {
			fmt.Fprintf(os.Stderr, "USAGE:\n\t%s <filename>\n", idiom_cmdln.Name)
			idiom_log.Tracef("[main] cmdln args: %#v", idiom_cmdln.Args)
			idiom_log.Inform("[main] EXIT ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
			os.Exit(1)
			return
		}

		arg1 := idiom_cmdln.Args[0]

		filename = arg1
	}
	idiom_log.Inform("[main] FILENAME:", filename)

	// The ‘file’ variable will contain a reference to the open file for the single file of idiom source code
	// whose name is contained in the ‘filename’ variable.
	var file *os.File
	{
		var err error

		file, err = os.Open(filename)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: could not open file %q: %s\n", filename, err)
			idiom_log.Debugf("[main] ERROR TYPE: %T", err)
			idiom_log.Debugf("[main] ERROR: %q", err)
			idiom_log.Inform("[main] EXIT ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
			os.Exit(1)
			return
		}
		defer file.Close()
	}
	idiom_log.Informf("[main] file %q was successfully opened.", filename)


	// The for-loop here reads in a single line of code, and then tries to interpret it.
	{
		parser := idiom_parser.New(file)

		var cmdNumber int64 = 0
		var result idiom_string.Type
		for {
			idiom_log.Debug("[main] [for] BEGIN ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")

			cmdNumber++

			tokens, err := parser.Readln()
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: could read command %d in file file %q: %s\n", cmdNumber, filename, err)
				idiom_log.Debugf("[main] [for] ERROR TYPE: %T", err)
				idiom_log.Debugf("[main] [for] ERROR: %q", err)
				idiom_log.Inform("[main] [for] EXIT ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
				os.Exit(1)
			}
			if 0 >= len(tokens) {
		/////////////// CONTINUE
				idiom_log.Debug("[main] [for] CONTINUE ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
				continue
			}

			if idiom_log.Tracing() {
				idiom_log.Trace("[main] [for] TOKENS BEGIN")
				for i, token := range tokens {
					idiom_log.Tracef("[main] [for] token[%d] = %#v", i, token)
				}
				idiom_log.Trace("[main] [for] TOKENS END")
			}

			token0 := tokens[0]
			if idiom_token.EOF() == token0 {
				idiom_log.Inform("[main] [for] EOF")
				idiom_log.Inform("[main] [for] EXIT ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
				os.Exit(0)
			}

			var command string
			{
				var err error
				command, err = token0.Return()
				if nil != err {
					fmt.Fprintf(os.Stderr, "ERROR: could get command №%d in ile %q: %s\n", cmdNumber, filename, err)
					idiom_log.Debugf("[main] [for] ERROR TYPE: %T", err)
					idiom_log.Debugf("[main] [for] ERROR: %q", err)
					idiom_log.Inform("[main] [for] EXIT ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
					os.Exit(1)
				}
			}
			idiom_log.Inform("[main] [for] COMMAND:", command)

			parameters := tokens[1:]

			if idiom_log.Debugging() {
				idiom_log.Debug("[main] [for] PARAMETERS BEGIN")
				for _, p := range parameters {
					idiom_log.Debugf("[main] [for] %#v", p)
				}
				idiom_log.Debug("[main] [for] PARAMETERS END")
			}

			var p []idiom_string.Type = make([]idiom_string.Type, len(parameters))
			for i, parameter := range parameters {
				value, err := parameter.Return()
				if nil != err {
					fmt.Fprintf(os.Stderr, "ERROR: could get parameter №%d for command №%d in file %q: %s\n", i, cmdNumber, filename, err)
					idiom_log.Debugf("[main] [for] ERROR TYPE: %T", err)
					idiom_log.Debugf("[main] [for] ERROR: %q", err)
					idiom_log.Inform("[main] [for] EXIT ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
					os.Exit(1)
				}

				p[i] = idiom_string.Something(value)
			}

			result = idiom_kernel.Stem.Run(command, p...)
			if result.IsError() {
				fmt.Fprintln(os.Stderr)
				fmt.Fprintln(os.Stderr, "\x1b[30;41m"+"ERROR"+"\x1b[0m")
				fmt.Fprintln(os.Stderr, "COMMAND NUMBER:", cmdNumber)
				fmt.Fprintln(os.Stderr, "COMMAND:", command)
				fmt.Fprintln(os.Stderr, "PARAMETERS:", parameters)
				fmt.Fprintln(os.Stderr, "ERROR MESSAGE:", result.Err())
				idiom_log.Inform("[main] [for] EXIT ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
				os.Exit(1)
				return
			}

			idiom_log.Debug("[main] [for] END ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
		}
	}

	idiom_log.Inform("[main] END ◢◤◢◤◢◤◢◤◢◤◢◤◢◤")
}
