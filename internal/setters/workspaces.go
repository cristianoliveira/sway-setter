package setters

import (
	"fmt"
)

// SetWorkspaces configures workspaces to the provided outputs
// and focuses the workspace that is marked as focused.
//
// Use: '--print' to print the commands that would be executed.
func SetWorkspaces(workspaces []SwayWorkspace) error {
	swaymsg, err := ConnectToSway()
	if err != nil {
		return err
	}

	if len(workspaces) == 0 {
		return fmt.Errorf("Error: no workspaces provided")
	}

	var focusedWorkspace *SwayWorkspace
	for _, workspace := range workspaces {
		err = swaymsg.MoveWorkspaceToOutput(
			workspace.Name,
			workspace.Output,
		)

		if err != nil {
			return err
		}

		if workspace.Focused {
			focusedWorkspace = &workspace
		}
	}

	// Ensure that focused workspace received takes focus
	if focusedWorkspace != nil {
		swaymsg.FocusWorkspace(focusedWorkspace.Name)
	}

	return nil
}
