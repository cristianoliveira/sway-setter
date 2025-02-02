package e2e

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/testutils"
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
