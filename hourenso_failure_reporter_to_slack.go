package failurereporter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type horensoOut struct {
	command     string
	commandArgs []string
	output      string
	stdout      string
	stderr      string
	exitCode    int
	result      string
	pid         int
	startAt     string
	endAt       string
	hostname    string
	systemTime  float32
	userTime    float32
}

func parseHorensoOut(stdin io.Reader) (*horensoOut, error) {
	ho := new(horensoOut)

	text, err := ioutil.ReadAll(stdin)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(text), ho)
	return ho, err
}

// Run the reporter
func Run(stdin io.Reader, stdout io.Writer, stderr io.Writer, c SlackClient) int {
	ho, err := parseHorensoOut(stdin)

	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 1
	}

	if ho.exitCode == 0 {
		return 0
	}

	c.Post(ho.stderr)

	return 0
}
