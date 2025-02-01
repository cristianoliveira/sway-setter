package setters

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/sway"
	"github.com/cristianoliveira/sway-setter/internal/testutils"
)

func TestOutputSetter(t *testing.T) {
	t.Run("SetOutputPosition", func(t *testing.T) {
		con := testutils.MockedConnector{}
		sway.SwayIPCConnector = &con

		swayOutputs := []SwayOutput{
			{
				Name: "eDP-1",
				Rect: &OutputRect{
					X:      0,
					Y:      0,
					Width:  1920,
					Height: 1080,
				},
			},
			{
				Name: "HDMI-A-1",
				Rect: &OutputRect{
					X:      1920,
					Y:      0,
					Width:  1920,
					Height: 1080,
				},
			},
		}

		expectedCommands := []string{
			"output eDP-1 position 0 0",
			"output HDMI-A-1 position 1920 0",
		}

		err := SetOutputs(swayOutputs)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if len(con.CommandsHistory) != len(expectedCommands) {
			t.Errorf(
				"Expected %d commands to be executed, got %d",
				len(expectedCommands),
				len(con.CommandsHistory),
			)
		}

		for i, command := range con.CommandsHistory {
			if command != expectedCommands[i] {
				t.Errorf("Expected: \n %s\nGot: %s", expectedCommands[i], command)
			}
		}
	})
}
