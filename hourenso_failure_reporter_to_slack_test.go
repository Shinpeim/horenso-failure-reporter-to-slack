package failurereporter

import (
	"bytes"
	"strings"
	"testing"
)

type mockedSlackClient struct {
	Text string
}
func (m *mockedSlackClient) Post(text string) error {
	m.Text = text
	return nil
}
func newMockedSlcakClient() *mockedSlackClient{
	return &mockedSlackClient{"initial text"}
}

func TestInvalidJson(t *testing.T) {
	mockedStdOut := &bytes.Buffer{}
	mockedStdErr := &bytes.Buffer{}
	invalidJSONReader := strings.NewReader("invalid json")
	mockedSlackClient := newMockedSlcakClient()

	exitCode := Run(invalidJSONReader, mockedStdOut, mockedStdErr, mockedSlackClient)

	if exitCode != 1 {
		t.Fatal("can't handle invalid JSON")
	}
	if mockedStdErr.String() == "" {
		t.Fatal("got no error message when given invalid json")
	}
}

func TestValidJson(t *testing.T) {
	json := `{
		"command": "command",
		"commandArgs": [
		  "command"
		],
		"output": "1",
		"stdout": "1",
		"stderr": "1",
		"exitCode": 0,
		"result": "command exited with code: 0",
		"pid": 95030,
		"startAt": "2015-12-28T00:37:10.494282399+09:00",
		"endAt": "2015-12-28T00:37:10.546466379+09:00",
		"hostname": "webserver.example.com",
		"systemTime": 0.034632,
		"userTime": 0.026523
	}`
	validJSONReader := strings.NewReader(json)

	_, err := parseHorensoOut(validJSONReader)

	if err != nil {
		t.Fatal("failed to parse json")
	}
}

func TestSucceededCommand(t *testing.T) {
	json := `{
		"command": "command",
		"commandArgs": [
		  "command"
		],
		"output": "1",
		"stdout": "1",
		"stderr": "1",
		"exitCode": 0,
		"result": "command exited with code: 0",
		"pid": 95030,
		"startAt": "2015-12-28T00:37:10.494282399+09:00",
		"endAt": "2015-12-28T00:37:10.546466379+09:00",
		"hostname": "webserver.example.com",
		"systemTime": 0.034632,
		"userTime": 0.026523
	}`

	mockedStdOut := &bytes.Buffer{}
	mockedStdErr := &bytes.Buffer{}
	jr := strings.NewReader(json)
	mockedSlackClient := newMockedSlcakClient()

	exitCode := Run(jr, mockedStdOut, mockedStdErr, mockedSlackClient)

	if exitCode != 0 {
		t.Fatal("failed to handle failed command")
	}

	if mockedSlackClient.Text != "initial text" {
		t.Fatal("slack client was called when the command succeeded")
	}

}