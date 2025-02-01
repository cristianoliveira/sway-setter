package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestSetOutputs(t *testing.T) {
	cases := []struct {
		name       string
		args       []string
		stdin      string
		expectFail bool
	}{
		{
			name: "set outputs with invalid output",
			args: []string{
				"-t", "set_outputs",
				"--print",
			},
			stdin: `
			[
				{
					"id":1,
					"name":"HDMI-A-0",
					"foobar":{
						"x":0,
						"y":0,
						"width":1920,
						"height":1080
					}
				}
			]`,
			expectFail: true,
		},
		{
			name: "set outputs without name",
			args: []string{
				"-t", "set_outputs",
				"--print",
			},
			stdin: `
			[
				{
					"id": 383,
					"foo": "DP-2",
					"type": "output",
					"orientation": "none",
					"rect": {
						"x": 1920,
						"y": 0,
						"width": 1920,
						"height": 1080
					}
				}
			]`,
			expectFail: true,
		},
		{
			name: "set outputs one output",
			args: []string{
				"-t", "set_outputs",
				"--print",
			},
			stdin: `
			[
				{
					"id": 383,
					"name": "DP-2",
					"type": "output",
					"orientation": "none",
					"rect": {
						"x": 1920,
						"y": 0,
						"width": 1920,
						"height": 1080
					}
				}
			]`,
			expectFail: false,
		},
		{
			name: "set outputs multiple outputs",
			args: []string{
				"-t", "set_outputs",
				"--print",
			},
			stdin: `
			[
				{
					"id": 383,
					"name": "DP-2",
					"type": "output",
					"orientation": "none",
					"rect": {
						"x": 1920,
						"y": 0,
						"width": 1920,
						"height": 1080
					}
				},
				{
					"id": 384,
					"name": "DP-1",
					"type": "output",
					"orientation": "none",
					"rect": {
						"x": 0,
						"y": 0,
						"width": 1920,
						"height": 1080
					}
				}
			]`,
			expectFail: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(tt *testing.T) {
			cmd := exec.Command("../bin/sway-setter", tc.args...)

			var out, errOut bytes.Buffer
			stdin := bytes.NewBufferString(tc.stdin)
			cmd.Stdin = stdin
			cmd.Stdout = &out
			cmd.Stderr = &errOut

			err := cmd.Run()
			if err != nil && !tc.expectFail {
				tt.Fatalf("\nUnexpected fail:\n%s\nstderr:%s\n stdout:%s", err, errOut.String(), out.String())
			}

			if tc.expectFail && err == nil {
				tt.Fatalf("\nExpected fail but got success")
			}

			var prettyJson bytes.Buffer
			err = json.Indent(&prettyJson, []byte(tc.stdin), "", "  ")
			if err != nil {
				tt.Fatalf("\nError indenting json: %s", err)
			}
			snaps.MatchSnapshot(
				tt,
				fmt.Sprintf("%s <<EOF\n%s\nEOF", cmd.String(), prettyJson.String()),
				out.String(),
			)
		})
	}
}
