
// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"bufio"
//	"fmt"
//	"io"
//	"io/ioutil"
	"log"
	"os"
	"regexp"
)

const DEBUG = false

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func fatalUsage() {
	log.Fatalf("usage: %v [input.go [output.go]]\npurpose: move braces around to avoid their loneliness", os.Args[0]);
}

/**
 * Given the name of exactly one readable file (that is presumed to be
 * golang source code), output that same file with any "lonely braces"
 * moved to the previous line, as is required by current golang syntax.
 *
 * This is my first golang program, please be kind.
 * 
 * Limitations:
 * (1) Assumes there is at least one line in the file.
 * (2) Does not work on empty files.
 *
 * Much copy/paste from:
 * https://gobyexample.com/line-filters
 * https://gobyexample.com/reading-files
 * https://gobyexample.com/command-line-arguments
 * 
 */
func main() {

	input:=os.Stdin
	output:=os.Stdout
	var err error;

	switch len(os.Args) {
		case 1 : {
			//no-args; input=os.Stdin, output=os.Stdout
		}
		case 2 : {
			//one argument; input specified (or request for help)
			if (os.Args[1][0]=='-') {
				fatalUsage()
			}

			input, err = os.Open(os.Args[1])
			check(err)
			defer input.Close()
		}
		case 3 : {
			//two arguments; input & output specified
			input, err = os.Open(os.Args[1])
			check(err)
			defer input.Close()

			output, err = os.OpenFile(os.Args[2], os.O_WRONLY  | os.O_CREATE | os.O_TRUNC, 0666)
			check(err)
			defer output.Close()
		}
		default:
			fatalUsage()
	}

	scanner := bufio.NewScanner(input);

	lonelyBrace := regexp.MustCompile("[:space:]*{[:space:]*$");
	lonelyElse := regexp.MustCompile("[:space:]*else[:space:]*$");

	if (DEBUG) { log.Println("ready to read") }

	if scanner.Scan() {
		var previous string = scanner.Text();
		if (DEBUG) { log.Println("initial read: ", previous) }

		for scanner.Scan() {
			next := scanner.Text();
			if (DEBUG) { log.Println("subsequent read: ", next) }
			if lonelyBrace.MatchString(next) {
				if (DEBUG) { log.Println("found lonely brace") }
				previous = previous + " {";
				//NB: "next" is forgotten (it contains the lonely brace).
				//NB: previous is not printed, in case we have multiple items to process (e.g. "} else {")
			} else if (lonelyElse.MatchString(next)) {
				if (DEBUG) { log.Println("found lonely 'else'") }
				previous = previous + " else";
			} else {
				if (DEBUG) { log.Println("printing: ", previous) }
				output.WriteString(previous);
				output.WriteString("\n");
				previous = next;
			}
		}

		if (DEBUG) { log.Println("printing: ", previous) }
		output.WriteString(previous);
		output.WriteString("\n");
	}

}


