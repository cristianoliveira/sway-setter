package setters

import (
	"fmt"
)

type OutputRect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Mode struct {
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	RefreshRate int    `json:"refresh_rate"`
	AspectRatio string `json:"picture_aspect_ratio"`
}

type SwayOutput struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Active     bool        `json:"active"`
	Dpms       bool        `json:"dpms"`
	Transform  string      `json:"transform"`
	Rect       *OutputRect `json:"rect"`
	CurentMode *Mode       `json:"current_mode"`
}

// toCommand returns the command to set the output
func (s SwayOutput) toCommand() (string, error) {
	if len(s.Name) == 0 {
		return "", fmt.Errorf("Error: output name is empty")
	}

	if s.Rect == nil {
		return "", fmt.Errorf("Error: output rect is empty")
	}

	command := fmt.Sprintf(
		"output %s position %d %d",
		s.Name,
		s.Rect.X,
		s.Rect.Y,
	)

	if s.CurentMode != nil {
		command += fmt.Sprintf(
			" resolution %dx%d@%dHz",
			s.CurentMode.Width,
			s.CurentMode.Height,
			s.CurentMode.RefreshRate,
		)
	}

	if len(s.Transform) > 0 {
		command += fmt.Sprintf(" transform %s", s.Transform)
	}

	return command, nil
}

// SetOutputs apply the given outputs configuration
// into the sway via swaymsg
// See 'man 5 sway-output' and 'man swaymsg'
func SetOutputs(outputs []SwayOutput) error {
	swaymsg, err := ConnectToSway()
	if err != nil {
		return err
	}

	if len(outputs) == 0 {
		return fmt.Errorf("Error: no outputs provided")
	}

	for _, output := range outputs {
		if output.Rect == nil {
			return fmt.Errorf("Error: output rect is empty")
		}

		command, err := output.toCommand()
		if err != nil {
			return err
		}

		err = swaymsg.Command(command)

		if err != nil {
			return err
		}
	}

	return nil
}
