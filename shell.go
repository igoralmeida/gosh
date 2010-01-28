package main

import (
	"fmt"
	"os"
)

//loop forever: get line from readch and treat it accordingly
func Gosh() {
	readch := make(chan string)
	exch := make(chan *command)

	initializeVars()
	initializePrompt()
	initializeCommands()
	parser := initializeParser()

	go readline(readch)
	go execute(exch)

	var line, outputStr string
	var status parseStatus
	var cmd *command
	var promptFunc (func() string)

	status = READY

	for {
		promptFunc = getPromptFunc(status)
		outputStr = promptFunc()
		fmt.Print(outputStr)

		line = <-readch

		cmd, status = parser.parse(line)

		switch status {
		case READY:
			if cmd != nil {
				//FIXME there's still some kind of
				// race condition happenning with pagers

				exch <- cmd
				_ = <-exch //block until execution releases
			}
			//case READING, SELECTING, TRACING:
		}
	}

}


func error(msg string) { fmt.Fprintln(os.Stderr, msg) }
