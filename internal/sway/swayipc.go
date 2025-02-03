package sway

import (
	"fmt"

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

func (c *Connector) Connect() (*SwayMsgConnection, error) {
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

type CustomExecutor struct {
	HandleExecute func(command string) ([]byte, error)
}

func (m CustomExecutor) Execute(command string) ([]byte, error) {
	return m.HandleExecute(command)
}

func (c CustomExecutor) Connect() (*SwayMsgConnection, error) {
	return &SwayMsgConnection{
		SwayIPC: &CustomExecutor{},
	}, nil
}

// StdOutputConnector is a connector that outputs the commands to the standard output.
// Used when running with --output flag.
type StdOutputConnector struct{}

func (c StdOutputConnector) Connect() (*SwayMsgConnection, error) {
	return &SwayMsgConnection{
		SwayIPC: &CustomExecutor{
			HandleExecute: func(command string) ([]byte, error) {
				fmt.Printf("%s\n", command)
				return []byte{}, nil
			},
		},
	}, nil
}

func ConfigStdoutConnector() {
	SwayIPCConnector = &StdOutputConnector{}
}

func ConnectToSway() (*SwayMsgConnection, error) {
	swaymsg, err := SwayIPCConnector.Connect()
	if err != nil {
		return nil, err
	}

	return swaymsg, nil
}
