package e2e

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/testutils"
	"github.com/gkampitakis/go-snaps/snaps"
)

func TestSetValidations(t *testing.T) {
	cases := []struct {
		name       string
		args       []string
		stdin      string
		expectFail bool
	}{
		{
			name: "set unkonwn workspaces",
			args: []string{
				"-t", "unknown",
				"--print",
			},
			stdin:      `[{"id":1,"name":"1","output":"HDMI-A-0","focused":true}]`,
			expectFail: true,
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
