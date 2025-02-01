package setters

import (
	"github.com/cristianoliveira/sway-setter/internal/sway"
)

type SwayWorkspace struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Output string `json:"output"`
	Focused bool  `json:"focused"`
}

func SetWorkspaces(workspaces []SwayWorkspace) {
	swaymsg, err := sway.SwayIPCConnector.Connect()
	if err != nil {
		panic(err)
	}

	var focusedWorkspace *SwayWorkspace
	for _, workspace := range workspaces {
		swaymsg.MoveWorkspaceToOutput(
			workspace.Name,
			workspace.Output,
		)

		if workspace.Focused {
			focusedWorkspace = &workspace
		}
	}

	// Ensure that focused workspace received takes focus
	if focusedWorkspace != nil {
		swaymsg.FocusWorkspace(focusedWorkspace.Name)
	}
}
