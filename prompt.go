
package main

import (
	"regexp"
	"os"
)

type promptParser struct {
	action (func() string)
	re     *regexp.Regexp
}

var promptSequences map[string]*promptParser

func initializePrompt() {
	var m = map[string]*promptParser{
		`\\u`: &promptParser{getUser, nil},
		`\\h`: &promptParser{getHost, nil},
		`\\w`: &promptParser{getCwd, nil},
		`\\\$`: &promptParser{getTermination, nil}, //return # if root, else $
		//`\\W`, //basename of \w with ~
		//`\\s`, //shell name
	}

	for seq, parser := range m {
		parser.re = regexp.MustCompile(seq)
	}

	promptSequences = m
}

func parsePromptOutput(str string) string {
	outStr := str

	for _, parser := range promptSequences {
		outStr = parser.re.ReplaceAllString(outStr, parser.action())
	}

	return outStr
}

func getPromptFunc(status parseStatus) (func() string) {
	var promptVar string

	switch status {
	case READY:
		promptVar = "PS1"
	case READING:
		promptVar = "PS2"
	case SELECTING:
		promptVar = "PS3"
	case TRACING:
		promptVar = "PS4"
	}
	return func() string { return parsePromptOutput(GetLocalVar(promptVar)) }
}

func getUser() string { return os.Getenv(`USER`) }
func getCwd() string {
	cwd, _ := os.Getwd()
	return cwd
}
func getHost() string {
	host, err := os.Hostname()

	if err != nil {
		error(err.String() + "hostname lookup failed")
		os.Exit(-1)
	}
	return host
}
func getTermination() string {
	if os.Geteuid() == 0 {
		return "#"
	}
	return "$"
}
