package main

import (
	"os"

	"github.com/Shinpeim/hourenso-failure-reporter-to-slack"
)

func main() {
	os.Exit(failurereporter.Run(os.Stdin, os.Stdout, os.Stderr, failurereporter.NewSlackClient()))
}
