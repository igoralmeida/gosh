package main

import (
	"os"
	"exec"
)

var LAST_RETURN_VALUE int

/* Execute command coming from inch channel */
func execute(inch chan *command) {
	var c *command
	var cmd *exec.Cmd
	var err os.Error

	var argv0 string
	var argv []string

	for {
		c = <-inch

		switch c.cmdtype {
		case BUILTIN:
			cb := BUILTIN_COMMANDS[c.name]
			LAST_RETURN_VALUE = cb(c.args)
		case EXTERNAL:
			if argv0, err = exec.LookPath(c.name); err != nil {
				goto Error
			}

			if len(c.args) > 0 {
				argv = c.args
			}
			envv := os.Environ()

			cmd, err = exec.Run(argv0, argv, envv,
				exec.PassThrough, exec.PassThrough, exec.PassThrough)

			err = cmd.Close()
			//TODO update LAST_RETURN_VALUE

		Error:
			if err != nil {
				error(err.String())
			}
		}
		inch <- nil //release blockage in main loop()
	}
}
