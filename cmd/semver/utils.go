package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func readInput(readFrom io.Reader) string {
	data, err := ioutil.ReadAll(readFrom)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getCLIInput(c *cli.Context) string {
	output := ""
	fi, _ := os.Stdin.Stat()
	if fi.Mode()&os.ModeCharDevice == 0 {
		output = readInput(os.Stdin)
	} else if len(c.Args()) > 0 {
		output = c.Args()[0]
	}
	return strings.Trim(output, "\n\r_- ")
}

func getCurrentWorkingDirectory() string {
	if workingDirectory, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		return workingDirectory
	}
}
