package e2e

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/testutils"
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
				"workspaces",
				"--print",
			},
			// Expect a json array
			stdin:      `{"id":1,"name":"1","output":"HDMI-A-0","focused":true}`,
			expectFail: true,
		},
		{
			name: "set workspaces with incomplete json",
			args: []string{
				"workspaces",
				"--print",
			},
			// Expect a json array
			stdin:      `{"id":1,"name":"1"`,
			expectFail: true,
		},
		{
			name: "set workspaces one workspace",
			args: []string{
				"workspaces",
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
				"workspaces",
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
			cmd := testutils.NewCommandTesting(tc.stdin, tc.args...)

			err := cmd.Run()
			if err != nil && !tc.expectFail {
				tt.Fatalf(err.Error())
			}

			command, err := cmd.String()
			if err != nil {
				tt.Fatalf(err.Error())
			}

			snaps.MatchSnapshot(tt, command)
		})
	}
}
