package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type CommandTesting struct {
	Input  string
	StdIn  io.Reader
	StdOut *bytes.Buffer
	StdErr *bytes.Buffer
	Cmd    *exec.Cmd
}

func NewCommandTesting(input string, args ...string) CommandTesting {
	cmd := exec.Command("go", append([]string{"run", "../."}, args...)...)

	var out, errOut bytes.Buffer
	stdin := bytes.NewBufferString(input)
	cmd.Stdin = stdin
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	return CommandTesting{
		Input:  input,
		StdIn:  stdin,
		StdOut: &out,
		StdErr: &errOut,
		Cmd:    cmd,
	}
}

func (c *CommandTesting) Run() error {
	return c.Cmd.Run()
}

func (c *CommandTesting) StdoutString() string {
	return c.StdOut.String()
}

func (c *CommandTesting) StderrString() string {
	return c.StdErr.String()
}

func (c *CommandTesting) HasFailed() bool {
	return c.Cmd.ProcessState.ExitCode() != 0
}

// String returns the command that would be executed on the terminal
// with the input provided. This is useful for documentation so snapshots
// reflects the real command and with its args, input and output.
func (c *CommandTesting) String() (string, error) {
	var prettyJson bytes.Buffer
	err := json.Indent(&prettyJson, []byte(c.Input), "", "  ")
	if err != nil {
		prettyJson.WriteString(c.Input)
	}

	// Replace "go run ." with the sway-setter
	args := strings.Join(c.Cmd.Args[3:], " ")
	return fmt.Sprintf(
		"%s %s <<EOF\n%s\nEOF\n%s%s",
		"sway-setter",
		args,
		prettyJson.String(),
		c.StdOut.String(),
		c.StdErr.String(),
	), nil
}
