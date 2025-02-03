package sway

import "fmt"

type SwayMsgResult struct {
	Success bool `json:"success"`
}

func (s SwayMsgConnection) Commands(commands *[]string) error {
	var errorMsg string
	for _, cmd := range *commands {
		_, err := s.SwayIPC.Execute(cmd)
		if err != nil {
			errorMsg += fmt.Sprintf("%s error: %s\n", cmd, err)
		}
	}

	if errorMsg != "" {
		return fmt.Errorf("One or more commands failed:\n%s", errorMsg)
	}

	return nil
}
