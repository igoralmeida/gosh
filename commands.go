package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type command struct {
	cmdtype int
	name    string
	args    []string
}

func (c *command) String() string {
	return fmt.Sprint("{", c.cmdtype, "-", c.name,
		"-", len(c.args), c.args, "}")
}

const (
	ALIAS = 1 << iota
	KEYWORD
	FUNCTION
	BUILTIN
	EXTERNAL
)

var BUILTIN_COMMANDS map[string](func([]string) int)

func initializeCommands() {
	var m = map[string](func([]string) int){
		"echo": echo,
		"exit": exit,
		"kill": kill,
		"pwd": pwd,
		"cd": cd,
	}
	BUILTIN_COMMANDS = m

}


func echo(args []string) int {
	fmt.Println(strings.Join(args[1:], " "))
	return 0
}
func exit(args []string) int {
	var code int = LAST_RETURN_VALUE
	var err os.Error

	switch {
	case len(args) > 2:
		fmt.Fprintln(os.Stderr, "exit: too many arguments")
	case len(args) > 1:
		code, err = strconv.Atoi(args[1])
	}

	if err != nil {
		code = 0 // FIXME 0 srsly?
	}

	os.Exit(code)
	return 0
}
func kill(args []string) int { return 0 }
func pwd(args []string) int {
	wd, err := os.Getwd()

	if err != nil {
		error(err.String() + "obscure error")
		return -1
	}

	fmt.Println(wd)
	return 0
}
func cd(args []string) int {
	var whereTo, cwd, old string

	whereTo = os.Getenv("HOME")
	cwd = os.Getenv("PWD")
	old = os.Getenv("OLDPWD")

	if len(cwd) == 0 {
		cwd, _ = os.Getwd()
	}

	switch {
	case len(args[1:]) > 0: //TODO add -L and -P flags
		whereTo = args[1]
	}

	if whereTo == "-" {
		whereTo = old
	}

	err := os.Chdir(whereTo)

	if err != nil {
		error(err.String())
		return -1
	}

	//FIXME check for usage of . and ..
	os.Setenv("OLDPWD", cwd) //at least we tried twice
	os.Setenv("PWD", whereTo)
	return 0

}
