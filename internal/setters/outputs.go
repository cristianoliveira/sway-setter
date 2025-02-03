package setters

import (
	"fmt"
)

// outputToCommand returns the command to set the output
func outputToCommand(output SwayOutput) (string, error) {
	if len(output.Name) == 0 {
		return "", fmt.Errorf("Error: output name is empty")
	}

	if output.Rect == nil {
		return "", fmt.Errorf("Error: output rect is empty")
	}

	command := fmt.Sprintf(
		"output %s position %d %d",
		output.Name,
		output.Rect.X,
		output.Rect.Y,
	)

	if output.CurentMode != nil {
		command += fmt.Sprintf(
			" resolution %dx%d@%dHz",
			output.CurentMode.Width,
			output.CurentMode.Height,
			output.CurentMode.RefreshRate,
		)
	}

	if len(output.Transform) > 0 {
		command += fmt.Sprintf(" transform %s", output.Transform)
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

		command, err := outputToCommand(output)
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
