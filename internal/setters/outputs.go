package setters

import (
	"fmt"

	"github.com/cristianoliveira/sway-setter/internal/sway"
)

type OutputRect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type SwayOutput struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Active    bool       `json:"active"`
	Dpms      bool       `json:"dpms"`
	Transform string     `json:"transform"`
	Rect      OutputRect `json:"rect"`
}

func SetOutputs(outputs []SwayOutput) error {
	swaymsg, err := sway.SwayIPCConnector.Connect()
	if err != nil {
		return err
	}

	if len(outputs) == 0 {
		return fmt.Errorf("Error: no outputs provided")
	}

	for _, output := range outputs {
		err = swaymsg.SetOutputPosition(
			output.Name,
			output.Rect.X,
			output.Rect.Y,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
