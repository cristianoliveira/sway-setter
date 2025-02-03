package parser

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/sway"
	"github.com/cristianoliveira/sway-setter/internal/testutils"
)

func TestWorkspaceSetter(t *testing.T) {
	t.Run("SetWorkspaces", func(tt *testing.T) {
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

		commands, err := SetWorkspacesCommand(swayWorkspaces)
		if err != nil {
			tt.Fatalf("Expected no error, got: %s", err)
		}

		if len(*commands) != len(expectedCommands) {
			tt.Errorf("Expected 4 commands to be executed, got %d", len(*commands))
		}

		for i, command := range *commands {
			if command != expectedCommands[i] {
				tt.Errorf("Expected: \n %s\nGot: %s", expectedCommands[i], command)
			}
		}
	})
}

func TestWorkspaceSetterValidations(t *testing.T) {
	cases := []struct {
		title      string
		workspaces []SwayWorkspace
		errorMsg   string
		connector  sway.SwayConnector
	}{
		{
			title:      "empty workspaces",
			workspaces: []SwayWorkspace{},
			errorMsg:   "Error: no workspaces provided",
			connector:  &testutils.DinamicMockedConnector{},
		},

		{
			title: "workspace name is empty",
			workspaces: []SwayWorkspace{
				{
					Id:      1,
					Output:  "HDMI-A-0",
					Focused: true,
				},
			},
			errorMsg:  "Error: workspace name is empty",
			connector: &testutils.DinamicMockedConnector{},
		},

		{
			title: "output name is empty",
			workspaces: []SwayWorkspace{
				{
					Id:      1,
					Name:    "1",
					Focused: true,
				},
			},
			errorMsg:  "Error: output name is empty",
			connector: &testutils.DinamicMockedConnector{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.title, func(tt *testing.T) {
			commands, err := SetWorkspacesCommand(tc.workspaces)
			if commands != nil {
				tt.Errorf("Expected no command, got %v", commands)
			}

			if err == nil {
				tt.Errorf("Expected error, got nil")
			} else {
				if err.Error() != tc.errorMsg {
					tt.Errorf("Expected error message: %s, got: %s", tc.errorMsg, err.Error())
				}
			}
		})
	}
}
