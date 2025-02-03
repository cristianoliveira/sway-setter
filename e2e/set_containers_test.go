package e2e

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/testutils"
	"github.com/gkampitakis/go-snaps/snaps"
)

func TestSetContainers(t *testing.T) {
	cases := []struct {
		name       string
		args       []string
		stdin      string
		expectFail bool
	}{
		{
			name: "set containers with invalid json",
			args: []string{
				"-t", "set_containers",
				"--print",
			},
			// Expect a json array
			stdin:      `{"id":1,"name":"1","output":"HDMI-A-0","focused":true}`,
			expectFail: true,
		},
		{
			name: "set containers with incomplete json",
			args: []string{
				"-t", "set_containers",
				"--print",
			},
			// Expect a json array
			stdin:      `{"id":1,"name":"1"`,
			expectFail: true,
		},
		{
			name: "set workspaces without containers",
			args: []string{
				"-t", "set_containers",
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
			expectFail: true,
		},
		{
			name: "set containers with multiple workspaces",
			args: []string{
				"-t", "set_containers",
				"--print",
			},
			stdin: `{
				"id":1,
				"type": "root",
				"nodes":[
					{
						"id":2,
						"type":"output",
						"nodes":[
							{
								"id":3,
								"type":"workspace",
								"name":"1",
								"nodes":[
									{
										"id":4,
										"type":"con",
										"name":"1",
										"app_id":"app1"
									},
									{
										"id":5,
										"type":"con",
										"name":"2",
										"window_properties": {
											"title": "title1"
										}
									}
								]
							},
							{
								"id":3,
								"type":"workspace",
								"name":"2",
								"nodes":[
									{
										"id":11,
										"type":"con",
										"name":"1",
										"app_id": null,
										"marks": ["setter:1"]
									},
									{
										"id":22,
										"type":"con",
										"name":"2",
										"window_properties": {
											"class": "class1"
										}
									}
								]
							}
						]
					}
				]
			}`,
			expectFail: false,
		},
		{
			name: "set containers with floating nodes",
			args: []string{
				"-t", "set_containers",
				"--print",
			},
			stdin: `{
				"id":1,
				"type": "root",
				"nodes":[
					{
						"id":2,
						"type":"output",
						"nodes":[
							{
								"id":3,
								"type":"workspace",
								"name":"1",
								"nodes":[
									{
										"id":4,
										"type":"con",
										"name":"1",
										"app_id":"app1"
									},
									{
										"id":5,
										"type":"con",
										"name":"2",
										"window_properties": {
											"title": "title1"
										}
									}
								],
								"floating_nodes": [
									{
										"id": 22,
										"type": "floating_con",
										"name": "floating1",
										"app_id": "fn1",
										"rect": {
											"x": 1940,
											"y": 1103,
											"width": 962,
											"height": 902
										}
									}
								]
							}
						]
					}
				]
			}`,
			expectFail: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(tt *testing.T) {
			cmd := testutils.NewCommandTesting(tc.stdin, tc.args...)

			err := cmd.Run()
			if err != nil && !tc.expectFail {
				tt.Fatalf("Expected no error, got: %s", cmd.StderrString())
			}

			command, err := cmd.String()
			if err != nil {
				tt.Fatalf("Expected no error, got: %s - %s", err, command)
			}

			snaps.MatchSnapshot(tt, command)
		})
	}
}
