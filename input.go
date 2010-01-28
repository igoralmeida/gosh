package main

import (
	"os"
	"bufio"
	"strings"
)

/* Read from stdin and pass it along to outch channel */
func readline(outch chan string) {
	var linestr string
	var err os.Error
	linereader := bufio.NewReader(os.Stdin)

	for {
		linestr, err = linereader.ReadString('\n')
		if err != nil {
			if err != os.EOF {
				error(err.String() + "Input error")
			}
			os.Exit(-1)
		}
		/* remove delimiter and trim space */
		outch <- strings.TrimSpace(linestr[0:len(linestr)-1])
	}
}
