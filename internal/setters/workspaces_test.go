package setters

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/sway"
)

type MockedConnector struct{}

var commandHistory []string

func (c MockedConnector) Connect() (*sway.SwayMsgConnection, error) {
	return &sway.SwayMsgConnection{
		SwayIPC: &sway.CustomExecutor{
			HandleExecute: func(command string) ([]byte, error) {
				commandHistory = append(commandHistory, command)
				return []byte{}, nil
			},
		},
	}, nil
}

func TestWorkspaceSetter(t *testing.T) {
	t.Run("SetWorkspaces", func(t *testing.T) {
		commandHistory = []string{}
		sway.SwayIPCConnector = &MockedConnector{}

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

		if len(commandHistory) != len(expectedCommands) {
			t.Errorf("Expected 4 commands to be executed, got %d", len(commandHistory))
		}

		for i, command := range commandHistory {
			if command != expectedCommands[i] {
				t.Errorf("Expected: \n %s\nGot: %s", expectedCommands[i], command)
			}
		}
	})
}
