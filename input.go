package main

import (
	"readline"
)

/* Read from stdin and pass it along to outch channel */
func prompter(outch chan string) {
	var linestr *string
	var eofStr = `exit`

	for {
		promptFunc := getPromptFunc(Status())
		outputStr := promptFunc()

		switch linestr = readline.ReadLine(&outputStr); true {

		case linestr == nil:
			linestr = &eofStr

		case *linestr != "":
			readline.AddHistory(*linestr)
		}

		outch <- *linestr
		_ = <-outch //wait till we can prompt again
	}
}
