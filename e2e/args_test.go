package e2e

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/testutils"
	"github.com/gkampitakis/go-snaps/snaps"
)

func TestFlags(t *testing.T) {
	cases := []struct {
		name       string
		args       []string
		expectFail bool
	}{
		{
			name:       "no flag",
			args:       []string{},
			expectFail: true,
		},
		{
			name: "help flag",
			args: []string{
				"-h",
			},
			expectFail: false,
		},
		{
			name: "help flag",
			args: []string{
				"--help",
			},
			expectFail: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(tt *testing.T) {
			cmd := testutils.NewCommandTesting("", tc.args...)

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
