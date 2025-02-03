package parser

import (
	"fmt"
)

// MoveWorkspaceToOutput moves the workspace to the output
// See in 'man 5 sway'
func moveWorkspaceToOutput(workspaceName string, outputName string) (string, error) {
	if len(workspaceName) == 0 {
		return "", fmt.Errorf("Error: workspace name is empty")
	}

	if len(outputName) == 0 {
		return "", fmt.Errorf("Error: output name is empty")
	}

	return fmt.Sprintf("workspace %s; move workspace to output %s", workspaceName, outputName), nil
}

// FocusWorkspace focuses the workspace by name
// See in 'man 5 sway'
func focusWorkspace(workspaceName string) string {
	return fmt.Sprintf("workspace %s", workspaceName)
}

// SetWorkspacesCommand configures workspaces to the provided outputs
// and focuses the workspace that is marked as focused.
//
// Use: '--print' to print the commands that would be executed.
func SetWorkspacesCommand(workspaces []SwayWorkspace) (*[]string, error) {
	var commands []string
	if len(workspaces) == 0 {
		return nil, fmt.Errorf("Error: no workspaces provided")
	}

	var focusedWorkspace *SwayWorkspace
	for _, workspace := range workspaces {
		cmd, err := moveWorkspaceToOutput(
			workspace.Name,
			workspace.Output,
		)

		if err != nil {
			return nil, err
		}

		if workspace.Focused {
			focusedWorkspace = &workspace
		}

		commands = append(commands, cmd)
	}

	// Ensure that focused workspace received takes focus
	if focusedWorkspace != nil {
		cmd := focusWorkspace(focusedWorkspace.Name)
		commands = append(commands, cmd)
	}

	return &commands, nil
}
