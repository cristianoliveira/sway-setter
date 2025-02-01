package setters

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/sway"
)

type MockedConnector struct{}

var commandHistory []string

type MockedCommandExecutor struct{}

func (m MockedCommandExecutor) Execute(command string) ([]byte, error) {
	commandHistory = append(commandHistory, command)
	return []byte{}, nil
}

func (c MockedConnector) Connect() (*sway.SwayMsgConnection, error) {
	return &sway.SwayMsgConnection{
		SwayConnection: &MockedCommandExecutor{},
	}, nil
}

func TestWorkspaceSetter(t *testing.T) {
	t.Run("SetWorkspaces", func(t *testing.T) {
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

		SetWorkspaces(swayWorkspaces)

		if len(commandHistory) != 4 {
			t.Errorf("Expected 4 commands to be executed, got %d", len(commandHistory))
		}

		if commandHistory[0] != "workspace 1; move workspace to output eDP-1" {
			t.Errorf("Expected command 1 to be 'workspace 1; move workspace to output eDP-1', got %s", commandHistory[0])
		}

		if commandHistory[1] != "workspace 2; move workspace to output eDP-1" {
			t.Errorf("Expected command 2 to be 'workspace 2; move workspace to output eDP-1', got %s", commandHistory[1])
		}

		if commandHistory[2] != "workspace 3; move workspace to output eDP-1" {
			t.Errorf("Expected command 3 to be 'workspace 3; move workspace to output eDP-1', got %s", commandHistory[2])
		}

		if commandHistory[3] != "workspace 2" {
			t.Errorf("Expected at the end to focus on 'workspace 2', got %s", commandHistory[3])
		}
	})
}
