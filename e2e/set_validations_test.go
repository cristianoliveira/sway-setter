package e2e

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

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

			snaps.MatchSnapshot(
				tt,
				fmt.Sprintf("%s < data.json", cmd.String()),
				out.String(),
			)
		})
	}
}
