# selpg Implementation by Go

In this project, I use Golang to implement a command-line application *selpg*.

## Introduction

selpg is a command-line application which allows the user to choose specific pages of one file to print. The only thing a user should do is to enter the start page and the end page of the input file, as well as other arguments to describe the format of the file or destination to process the output. To see more details, please visit [开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html) .

## Usage

The command of selpg are like this:

```shell
USAGE: ./selpg -s start_page -e end_page [ -f=true|false | -l lines_per_page ] [ -d dest ] [ in_filename ]

-s int
	the starting page of your file (mandatory) (default -1)
-e int
	the ending page of your file (mandatory) (default -1)
-f bool
	if true, selpg will use '\f' to separate pages (optional) (default false)
-l int
	specify the number of lines in one page (optional) (default 72)
-d string
	specify the destination, usually another program or command, to process the output. It     is just like a hand-writing channel (optional) (default stdout)
in_filename string
	the name of the input file (optional) (default stdin)
```

## Design Details and Problems

The selpg application works in following order:

```shell
step 1: reads arguments from the command
step 2: recongnizes the arguments and validate them
step 3: acts what those arguments say
step 4: show the output
```

### `flag`

To implement step 1-4, we first need to create a `selpgArgs struct` to save arguments:

```go
type selpgArgs struct {
	startPage int // start page of the article
	endPage   int // end page of the article
	pageLen   int /* number of lines in one page, default value is 72,
	   can be overriden by "-l number" on command line */
	pageSeperator bool   // if true, seperate pages by \f
	printDest     string // destination of result pages
	inFilename    string // name of the file to be read
}
```

Next, we should use the standard library `flag` to help us to read and parse arguments from command. By using `flag` , we can recognize each argument and set their default value. However, the flag parser only supports for three format of inputting arguments:

```shell
-flag
-flag=x
-flag x  // non-boolean flags only
```

This cause variant format between different commands, and we should deal with this.

Another problem is that it doesn't support "mutual exclusive" arguments. For example, in selpg program, the ‘-l’ and '-f' arguments are mutual exclusive, but there is no function in `flag` can detect this error and we must process the situation by ourselves. In my selpg program, I suppose that ‘-f’ is prior, meaning that when ‘-f’ and '-l' appear simultaneously, the selpg ignores '-l' argument and executes only the function of '-f' argument. 

### `Cmd`

Another point I want to mentioned is the '-d' function, the hand-writing channel. We should use the `Cmd` struct of `os/exec` standard library to define a shell command and execute it in Go. We can use the `Cmd.Stdin` and `Cmd.Stdout` to imitate channel in linux os.   The `Cmd.Stdin` defines the input of the command and the `Cmd.Stdout` defines the output of the command. For example, By setting `Cmd.Stdin = string1` and `Cmd.Stdout = os.Stdout` , we can see the output of the command dealing with `string1` on the screen.

## Examples

The file used for test is like this, which has 30 pages, 72 lines per page, and there is a character '\f' at the end of each page:

```shell
$ cat inputFile
No. 1 line of No.1 page
No. 2 line of No.1 page
No. 3 line of No.1 page
No. 4 line of No.1 page
No. 5 line of No.1 page
No. 6 line of No.1 page
No. 7 line of No.1 page
No. 8 line of No.1 page
No. 9 line of No.1 page
No. 10 line of No.1 page
No. 11 line of No.1 page
No. 12 line of No.1 page
......
No. 60 line of No.30 page
No. 61 line of No.30 page
No. 62 line of No.30 page
No. 63 line of No.30 page
No. 64 line of No.30 page
No. 65 line of No.30 page
No. 66 line of No.30 page
No. 67 line of No.30 page
No. 68 line of No.30 page
No. 69 line of No.30 page
No. 70 line of No.30 page
No. 71 line of No.30 page
No. 72 line of No.30 page
```

Read from file:

```shell
./selpg -s 1 -e 1 inputFile
No. 1 line of No.1 page
No. 2 line of No.1 page
No. 3 line of No.1 page
No. 4 line of No.1 page
No. 5 line of No.1 page
No. 6 line of No.1 page
No. 7 line of No.1 page
No. 8 line of No.1 page
No. 9 line of No.1 page
No. 10 line of No.1 page
No. 11 line of No.1 page
No. 12 line of No.1 page
......
No. 72 line of No.1 page
./selpg: done
```

Read from stdin:

```shell
./selpg -s 1 -e 1 < inputFile
No. 1 line of No.1 page
No. 2 line of No.1 page
No. 3 line of No.1 page
No. 4 line of No.1 page
No. 5 line of No.1 page
No. 6 line of No.1 page
No. 7 line of No.1 page
No. 8 line of No.1 page
No. 9 line of No.1 page
No. 10 line of No.1 page
No. 11 line of No.1 page
No. 12 line of No.1 page
......
No. 72 line of No.1 page
./selpg: done
```

Read from output of other command:

```shell
ls -a | ./selpg -s 1 -e 1
.
..
inputFile
selpg
selpg.go
./selpg: done
```

Save output to another file:

```shell
./selpg -s 10 -e 20 inputFile >outputFile
./selpg: done
```

Save error message to another file: (note: no "selpg: done" to be displayed on the screen)

```shell
./selpg -s 10 -e 20 inputFile 2>errorFile

No. 1 line of No.10 page
No. 2 line of No.10 page
No. 3 line of No.10 page
No. 4 line of No.10 page
No. 5 line of No.10 page
No. 6 line of No.10 page
No. 7 line of No.10 page
No. 8 line of No.10 page
No. 9 line of No.10 page
No. 10 line of No.10 page
No. 11 line of No.10 page
No. 12 line of No.10 page
......
No. 70 line of No.20 page
No. 71 line of No.20 page
No. 72 line of No.20 page
```

Specify the page length:

```shell
./selpg -s 10 -e 20 -l 2 inputFile
No. 19 line of No.1 page
No. 20 line of No.1 page
No. 21 line of No.1 page
No. 22 line of No.1 page
No. 23 line of No.1 page
No. 24 line of No.1 page
No. 25 line of No.1 page
No. 26 line of No.1 page
No. 27 line of No.1 page
No. 28 line of No.1 page
No. 29 line of No.1 page
No. 30 line of No.1 page
No. 31 line of No.1 page
No. 32 line of No.1 page
No. 33 line of No.1 page
No. 34 line of No.1 page
No. 35 line of No.1 page
No. 36 line of No.1 page
No. 37 line of No.1 page
No. 38 line of No.1 page
No. 39 line of No.1 page
No. 40 line of No.1 page
./selpg: done
```

Use '\f' to separate the pages:

```shell
./selpg -s 1 -e 2 -f inputFile
No. 1 line of No.1 page
No. 2 line of No.1 page
No. 3 line of No.1 page
No. 4 line of No.1 page
No. 5 line of No.1 page
No. 6 line of No.1 page
No. 7 line of No.1 page
No. 8 line of No.1 page
No. 9 line of No.1 page
No. 10 line of No.1 page
No. 11 line of No.1 page
No. 12 line of No.1 page
......
No. 70 line of No.2 page
No. 71 line of No.2 page
No. 72 line of No.2 page
./selpg: done
```

Use hand-writing channel:

```shell
./selpg -s 1 -e 2 -d "sort -r" inputFile
No. 9 line of No.2 page
No. 9 line of No.1 page
No. 8 line of No.2 page
No. 8 line of No.1 page
No. 7 line of No.2 page
No. 7 line of No.1 page
No. 72 line of No.2 page
No. 72 line of No.1 page
No. 71 line of No.2 page
No. 71 line of No.1 page
No. 70 line of No.2 page
No. 70 line of No.1 page
No. 6 line of No.2 page
No. 6 line of No.1 page
......
No. 14 line of No.2 page
No. 14 line of No.1 page
No. 13 line of No.2 page
No. 13 line of No.1 page
No. 12 line of No.2 page
No. 12 line of No.1 page
No. 11 line of No.2 page
No. 11 line of No.1 page
No. 10 line of No.2 page
No. 10 line of No.1 page
./selpg: done
```