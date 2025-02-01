package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestSetWorkspaces(t *testing.T) {
	cases := []struct {
		name       string
		args       []string
		stdin      string
		expectFail bool
	}{
		{
			name: "set workspaces with invalid json",
			args: []string{
				"-t", "set_workspaces",
				"--print",
			},
			// Expect a json array
			stdin:      `{"id":1,"name":"1","output":"HDMI-A-0","focused":true}`,
			expectFail: true,
		},
		{
			name: "set workspaces with incomplete json",
			args: []string{
				"-t", "set_workspaces",
				"--print",
			},
			// Expect a json array
			stdin:      `{"id":1,"name":"1"`,
			expectFail: true,
		},
		{
			name: "set workspaces one workspace",
			args: []string{
				"-t", "set_workspaces",
				"--print",
			},
			stdin: `[
				{
					"id":1,
					"name":"1",
					"output":"HDMI-A-0",
					"focused":true
				}
			]`,
			expectFail: false,
		},
		{
			name: "set workspaces with multiple workspaces",
			args: []string{
				"-t", "set_workspaces",
				"--print",
			},
			stdin: `[
				{
					"id":1,
					"name":"1",
					"output":"HDMI-A-0",
					"focused":false
				},
				{
					"id":2,
					"name":"2",
					"output":"DP-1",
					"focused":true
				},
				{
					"id":3,
					"name":"3",
					"output":"eDP-1",
					"focused":false
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

			if err := cmd.Run(); err != nil && !tc.expectFail {
				tt.Fatalf("\nUnexpected fail:\n%s\nstderr:%s\n stdout:%s", err, errOut.String(), out.String())
			}

			var prettyJson bytes.Buffer
			err := json.Indent(&prettyJson, []byte(tc.stdin), "", "  ")
			if err != nil {
				// Fall back to the original input if the json is invalid
				prettyJson.WriteString(tc.stdin)
			}

			snaps.MatchSnapshot(
				tt,
				fmt.Sprintf("%s<<EOF\n%s\nEOF", cmd.String(), prettyJson.String()),
				out.String(),
			)
		})
	}
}
