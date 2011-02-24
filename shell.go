package main

import (
	"fmt"
	"os"
)

var status parseStatus

//loop forever: get line from readch and treat it accordingly
func Gosh() {
	readch := make(chan string)
	exch := make(chan *command)

	initializeVars()
	initializePrompt()
	initializeCommands()
	parser := initializeParser()

	go prompter(readch)
	go execute(exch)

	var line string
	var cmd *command

	status = READY

	for {
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
			readch <- "" //release prompt
			//case READING, SELECTING, TRACING:
		}
	}

}

func Status() parseStatus {
	return status
}

func error(msg string) { fmt.Fprintln(os.Stderr, msg) }
func debug(msg string) { fmt.Fprint(os.Stdout, "DEBUG "); fmt.Fprintln(os.Stdout, msg) }

