package main

import (
	"strings"
)

type parseStatus int

type CommandParser struct {
	status parseStatus
}

const (
	READY parseStatus = 1 << iota
	READING
	SELECTING
	TRACING
)

var parser CommandParser

func initializeParser() (parser *CommandParser) {
	parser = &CommandParser{READY}

	return parser
}

/* TODO use some lex in the future instead:
 * check for redirection, etc.
 */
func (p *CommandParser) parse(line string) (action *command, status parseStatus) {
	var tmp []string

	tmp = strings.Split(line, " ", -1)

	if len(tmp) == 0 {
		tmp = []string{line}
	}

	cmdname, args := tmp[0], tmp

	/* default */
	cmdType := EXTERNAL
	status = READY

	_, ok := BUILTIN_COMMANDS[cmdname]

	if ok {
		cmdType = BUILTIN
		status = READY
	}

	action = nil
	if len(cmdname) > 0 {
		action = &command{cmdType, cmdname, args}
	}

	return action, status
}
