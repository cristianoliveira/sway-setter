package sway

import (
	"encoding/json"
	"fmt"
)

type SwayMsgResult struct {
	Success bool `json:"success"`
}

type SwayMsgr interface {
	Commmand(command string) error
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
	data, err := s.SwayConnection.Execute(command)
	if err != nil {
		panic(err)
	}

	result := []SwayMsgResult{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	if !result[0].Success {
		return fmt.Errorf("Command %s failed", command)
	}

	return nil
}
