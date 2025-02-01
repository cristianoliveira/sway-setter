package testutils

import (
	"github.com/cristianoliveira/sway-setter/internal/sway"
)

type MockedConnector struct{
	CommandsHistory []string
}

func (c *MockedConnector) Connect() (*sway.SwayMsgConnection, error) {
	return &sway.SwayMsgConnection{
		SwayIPC: &sway.CustomExecutor{
			HandleExecute: func(command string) ([]byte, error) {
				c.CommandsHistory = append(c.CommandsHistory, command)
				return []byte{}, nil
			},
		},
	}, nil
}

