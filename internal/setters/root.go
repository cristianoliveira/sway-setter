package setters

import "github.com/cristianoliveira/sway-setter/internal/sway"

func ConfigStdoutConnector() {
	sway.SwayIPCConnector = &sway.StdOutputConnector{}
}

func ConnectToSway() (*sway.SwayMsgConnection, error) {
	swaymsg, err := sway.SwayIPCConnector.Connect()
	if err != nil {
		return nil, err
	}

	return swaymsg, nil
}
