package setters

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/sway"
	"github.com/cristianoliveira/sway-setter/internal/testutils"
)

func TestWorkspaceSetter(t *testing.T) {
	t.Run("SetWorkspaces", func(t *testing.T) {
		con := testutils.MockedConnector{}
		sway.SwayIPCConnector = &con

		swayWorkspaces := []SwayWorkspace{
			{
				Name:    "1",
				Output:  "eDP-1",
				Focused: false,
			},
			{
				Name:    "2",
				Output:  "eDP-1",
				Focused: true,
			},
			{
				Name:    "3",
				Output:  "eDP-1",
				Focused: false,
			},
		}

		expectedCommands := []string{
			"workspace 1; move workspace to output eDP-1",
			"workspace 2; move workspace to output eDP-1",
			"workspace 3; move workspace to output eDP-1",
			"workspace 2",
		}

		SetWorkspaces(swayWorkspaces)

		if len(con.CommandsHistory) != len(expectedCommands) {
			t.Errorf("Expected 4 commands to be executed, got %d", len(con.CommandsHistory))
		}

		for i, command := range con.CommandsHistory {
			if command != expectedCommands[i] {
				t.Errorf("Expected: \n %s\nGot: %s", expectedCommands[i], command)
			}
		}
	})
}
