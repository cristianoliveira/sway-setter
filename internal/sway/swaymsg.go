package sway

import "strings"

type SwayMsgResult struct {
	Success bool `json:"success"`
}

func (s SwayMsgConnection) Commands(commands *[]string) error {
	command := strings.Join(*commands, "\n")

	_, err := s.SwayIPC.Execute(command)
	if err != nil {
		return err
	}

	return nil
}
