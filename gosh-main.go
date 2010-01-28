package main

import (
	"fmt"
	"flag"
	"strings"
)

/* Flags:

-i Interactive
-c Command string

*/

var (
	isInteractive    *bool = flag.Bool("i", true, "Specify an interactive shell")
	useCommandString *bool = flag.Bool("c", false, "Run commands from the command string")
)

func main() {
	flag.Parse()

	if *useCommandString {
		commandString := strings.Join(flag.Args(), " ")
		fmt.Println(commandString)
	} else if *isInteractive {
		Gosh()
	} else {
		flag.Usage()
	}
}
