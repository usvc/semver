package utils

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli"
)

func GetCLIInput(c *cli.Context) string {
	output := ""
	fi, _ := os.Stdin.Stat()
	if fi.Mode()&os.ModeCharDevice == 0 {
		output = readInput(os.Stdin)
	} else if len(c.Args()) > 0 {
		output = c.Args()[0]
	}
	return strings.Trim(output, "\n\r_- ")
}

func GetCurrentWorkingDirectory() string {
	if workingDirectory, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		return workingDirectory
	}
}

func ResolveAbsolutePath(providedPath string) string {
	if path.IsAbs(providedPath) {
		return providedPath
	}
	workingDirectory := GetCurrentWorkingDirectory()
	resolvedPath := path.Join(workingDirectory, providedPath)
	return resolvedPath
}

func readInput(readFrom io.Reader) string {
	data, err := ioutil.ReadAll(readFrom)
	if err != nil {
		panic(err)
	}
	return string(data)
}
