include $(GOROOT)/src/Make.$(GOARCH)

TARG=gosh
GOFILES=\
	gosh-main.go\
	shell.go\
	variables.go\
	prompt.go\
	parser.go\
	execution.go\
	commands.go\
	input.go\

include $(GOROOT)/src/Make.cmd

