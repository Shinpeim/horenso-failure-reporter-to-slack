package main

import (
	"os"

	"github.com/Shinpeim/horenso-failure-reporter-to-slack"
)

func main() {
	os.Exit(failurereporter.Run(os.Stdin, os.Stdout, os.Stderr, failurereporter.NewSlackClient()))
}
