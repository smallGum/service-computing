package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type selpgArgs struct {
	startPage int // start page of the article
	endPage   int // end page of the article
	pageLen   int /* number of lines in one page, default value is 72,
	   can be overriden by "-l number" on command line */
	pageSeperator bool   // if true, seperate pages by \f
	printDest     string // destination of result pages
	inFilename    string // name of the file to be read
}

var progname = "./selpg" // name of program, used to display error message

/**
* diplay the error and usage of selpg
* @param err the error message
 */
func printError(err string) {
	fmt.Fprintf(os.Stderr, err+"\n"+
		"\nUSAGE: %s -s start_page -e end_page [ -f=true|false | -l lines_per_page ] [ -d dest ] [ in_filename ]\n", progname)
	os.Exit(1)
}

/**
* split the arguments of command to initial the selpgArgs instance
* @param saAddr a pointer to a selpgArgs instance
 */
func initSelpgArgs(saAddr *selpgArgs) {
	flag.IntVar(&(saAddr.startPage), "s", -1, "start page of your file.  must greater than 0.")
	flag.IntVar(&(saAddr.endPage), "e", -1, "end page of your file.  must greater than or equal to start page.")
	flag.IntVar(&(saAddr.pageLen), "l", 72, "number of lines in one page. must greater than 0. default value is 30")
	flag.BoolVar(&(saAddr.pageSeperator), "f", false, "use [-f=true] to seperate pages by \\f")
	flag.StringVar(&(saAddr.printDest), "d", "", "specify destionation of output. default destination is stdout")

	flag.Parse()

	// check if the number of input file is valid
	if len(flag.Args()) > 1 {
		printError("cannot support to read more that one file immediately!")
	}
	if len(flag.Args()) == 0 {
		saAddr.inFilename = ""
	} else {
		saAddr.inFilename = flag.Args()[0]
	}

	// check if the the start page, end page and pageLen are entered correctly
	if saAddr.startPage < 1 {
		printError("start page must be set to be greater than 0!")
	}
	if saAddr.endPage < saAddr.startPage {
		printError("end page must be set to be grater than or equal to start page!")
	}
	if saAddr.pageLen < 1 {
		printError("page length must be greater than 0!")
	}
}

/**
* do operation correspoding to the arguments of command
 */
func runCommand() {
	var args selpgArgs
	initSelpgArgs(&args)

	// set the input source
	fin := os.Stdin
	var err error
	if args.inFilename != "" {
		fin, err = os.Open(args.inFilename)
		if err != nil {
			printError("could not open input file \"" + args.inFilename + "\"!")
		}
	}

	// set the output source
	fout := os.Stdout
	var cmd *exec.Cmd
	if args.printDest != "" {
		tmpStr := fmt.Sprintf("%s", args.printDest)
		cmd = exec.Command("sh", "-c", tmpStr)
		if err != nil {
			printError("could not open pipe to \"" + tmpStr + "\"!")
		}
	}

	// dealing with the page type
	var line string
	pageCnt := 1
	inputReader := bufio.NewReader(fin)
	rst := ""
	if args.pageSeperator == false {
		lineCnt := 0

		for true {
			line, err = inputReader.ReadString('\n')
			if err != nil { // error or EOF
				break
			}
			lineCnt++
			if lineCnt > args.pageLen {
				pageCnt++
				lineCnt = 1
			}
			if pageCnt >= args.startPage && pageCnt <= args.endPage {
				if args.printDest == "" {
					fmt.Fprintf(fout, "%s", line)
				} else {
					rst += line
				}
			}
		}
	} else {
		for true {
			c, _, erro := inputReader.ReadRune()
			if erro != nil { // error or EOF
				break
			}
			if c == '\f' {
				pageCnt++
			}
			if pageCnt >= args.startPage && pageCnt <= args.endPage {
				if args.printDest == "" {
					fmt.Fprintf(fout, "%c", c)
				} else {
					rst += string(c)
				}
			}
		}
	}

	if args.printDest != "" {
		cmd.Stdin = strings.NewReader(rst)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			printError("print error!")
		}
	}
	// end of dealing page type
	// end of setting output source

	if pageCnt < args.startPage {
		fmt.Fprintf(os.Stderr, "%s: start_page (%d) greater than total pages (%d), no output written\n", progname, args.startPage, pageCnt)
	} else {
		if pageCnt < args.endPage {
			fmt.Fprintf(os.Stderr, "%s: end_page (%d) greater than total pages (%d), less output than expected\n", progname, args.endPage, pageCnt)
		}
	}

	fin.Close()
	fout.Close()
	fmt.Fprintf(os.Stderr, "%s: done\n", progname)
}

func main() {
	runCommand()
}
