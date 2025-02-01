package sway

import (
	"fmt"
)

type SwayMsgResult struct {
	Success bool `json:"success"`
}

type SwayMsgr interface {
	MoveWorkspaceToOutput(workspaceName string, outputName string) error
	FocusWorkspace(workspaceName string) error
}

func (s SwayMsgConnection) MoveWorkspaceToOutput(workspaceName string, outputName string) error {
	command := fmt.Sprintf("workspace %s; move workspace to output %s", workspaceName, outputName)
	return s.Command(command)
}

func (s SwayMsgConnection) FocusWorkspace(workspaceName string) error {
	command := fmt.Sprintf("workspace %s", workspaceName)
	return s.Command(command)
}

func (s SwayMsgConnection) Command(command string) error {
	_, err := s.SwayIPC.Execute(command)
	if err != nil {
		return err
	}

	return nil
}
