package sway

import (
	"github.com/Difrex/gosway/ipc"
)

type SwayConnector interface {
	Connect() (*SwayMsgConnection, error)
}

type SwayMsgConnection struct {
	SwayIPC SwayIPCConnection
}

type Connector struct{}

type SwayIPCConnection interface {
	Execute(command string) ([]byte, error)
}

type SwayConnection struct {
	Con *ipc.SwayConnection
}

func (s *SwayConnection) Execute(command string) ([]byte, error) {
	data, err := s.Con.RunSwayCommand(command)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c Connector) Connect() (*SwayMsgConnection, error) {
	sc, err := ipc.NewSwayConnection()
	if err != nil {
		return nil, err
	}

	conn := &SwayConnection{
		Con: sc,
	}

	return &SwayMsgConnection{
		SwayIPC: conn,
	}, nil
}

var SwayIPCConnector SwayConnector = &Connector{}
