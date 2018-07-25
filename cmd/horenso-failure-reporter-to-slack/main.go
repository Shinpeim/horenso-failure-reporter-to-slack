package main

import (
	"os"

	"github.com/Shinpeim/horenso-failure-reporter-to-slack"
	"github.com/Shinpeim/horenso-failure-reporter-to-slack/slack-client"
)

func main() {
	os.Exit(failurereporter.Run(os.Stdin, os.Stdout, os.Stderr, slackclient.NewClient()))
}
