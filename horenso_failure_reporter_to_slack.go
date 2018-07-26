package failurereporter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type horensoOut struct {
	Command     string `json:"command"`
	CommandArgs []string `json:"commandArgs"`
	Output      string `json:"output"`
	Stdout      string `json:"stdout"`
	Stderr      string `json:"stderr"`
	ExitCode    int `json:"exitCode"`
	Result      string `json:"result"`
	Pid         int `json:"pid"`
	StartAt     string `json:"startAt"`
	EndAt       string `json:"endAt"`
	Hostname    string `json:"hostName"`
	SystemTime  float32 `json:"systemTime"`
	UserTime    float32 `json:"userTime"`
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

	if ho.ExitCode == 0 {
		return 0
	}

	err = c.Post(ho)
	if (err != nil) {
		fmt.Fprintln(stderr, err.Error())
		return 2
	}

	return 0
}
