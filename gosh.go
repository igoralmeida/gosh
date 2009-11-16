package gosh

import (
	"fmt";
	"os";
	"bufio";
	//"exec";
)

//loop forever: get line from readch and treat it accordingly:
// TODO
// is it a program call? use exec package
// is it a built-in command like 'cd'? do what it is supposed to do
func Gosh() {
	readch := make(chan string);
	//exch := make(chan string);

	go readline(readch);
	//go execute(exch);
	
	var command string;

	for {
		command = <-readch;
		fmt.Println(command)
	}

}

/* Read from stdin and pass it along to outch channel */
func readline(outch chan string) {
	var linestr string;
	var err os.Error;
	linereader := bufio.NewReader(os.Stdin);

	for {
		linestr, err = linereader.ReadString('\n');
		if err != nil {
			fmt.Println("ERROR");
			os.Exit(-1);
		}
		outch <- linestr;
	}
}

/* Execute command coming from inch channel */
func execute(inch chan string) {
}
