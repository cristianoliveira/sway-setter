package e2e

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

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
		testName := fmt.Sprintf("%s: %+v", tc.name, tc.args)
		t.Run(testName, func(tt *testing.T) {
			cmd := exec.Command("../bin/sway-setter", tc.args...)

			var out, errOut bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &errOut

			if err := cmd.Run(); err != nil && !tc.expectFail {
				tt.Fatalf("\nUnexpected fail:\n%s\nstderr:%s\n stdout:%s", err, errOut.String(), out.String())
			}

			snaps.MatchSnapshot(tt, cmd.String(), out.String())
		})
	}
}
