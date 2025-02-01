package sway

import (
	"github.com/Difrex/gosway/ipc"
)

type SwayConnector interface {
	Connect() (*SwayMsgConnection, error)
}

type SwayMsgConnection struct {
	SwayConnection SwayCommandExecutor
}

type Connector struct{}

type SwayCommandExecutor interface {
	Execute(command string) ([]byte, error)
}

type SwayCommand struct {
	Con *ipc.SwayConnection
}

func (s *SwayCommand) Execute(command string) ([]byte, error){
	data, err := s.Con.RunSwayCommand(command)
	return data, err
}

func (c Connector) Connect() (*SwayMsgConnection, error) {
	sc, err := ipc.NewSwayConnection()
	if err != nil {
		return nil, err
	}

	conn := &SwayCommand{
		Con: sc,
	}

	return &SwayMsgConnection{
		SwayConnection: conn,
	}, nil
}

var SwayIPCConnector SwayConnector = &Connector{}
