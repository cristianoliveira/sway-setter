package sway

import (
	"fmt"
)

type SwayMsgResult struct {
	Success bool `json:"success"`
}

func (s SwayMsgConnection) Command(command string) error {
	_, err := s.SwayIPC.Execute(command)
	if err != nil {
		return err
	}

	return nil
}

type SwayMsgr interface {
	MoveWorkspaceToOutput(workspaceName string, outputName string) error
	FocusWorkspace(workspaceName string) error
}

// MoveWorkspaceToOutput moves the workspace to the output
// See in 'man 5 sway'
func (s SwayMsgConnection) MoveWorkspaceToOutput(workspaceName string, outputName string) error {
	if len(workspaceName) == 0 {
		return fmt.Errorf("Error: workspace name is empty")
	}

	if len(outputName) == 0 {
		return fmt.Errorf("Error: output name is empty")
	}

	command := fmt.Sprintf("workspace %s; move workspace to output %s", workspaceName, outputName)
	return s.Command(command)
}

// FocusWorkspace focuses the workspace by name
// See in 'man 5 sway'
func (s SwayMsgConnection) FocusWorkspace(workspaceName string) error {
	command := fmt.Sprintf("workspace %s", workspaceName)
	return s.Command(command)
}
