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

				CurentMode: &Mode{
					Width:       1920,
					Height:      1080,
					RefreshRate: 60,
				},
			},
			{
				Name: "HDMI-A-1",

				Rect: &OutputRect{
					X:      1920,
					Y:      0,
					Width:  1280,
					Height: 1024,
				},

				CurentMode: &Mode{
					Width:       1280,
					Height:      1024,
					RefreshRate: 50,
				},
			},
			{
				Name: "HDMI-A-2",

				Rect: &OutputRect{
					X:      1920,
					Y:      0,
					Width:  1280,
					Height: 1024,
				},
			},
		}

		expectedCommands := []string{
			"output eDP-1 position 0 0 resolution 1920x1080@60Hz",
			"output HDMI-A-1 position 1920 0 resolution 1280x1024@50Hz",
			"output HDMI-A-2 position 1920 0",
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
				t.Errorf("Command:%d\nExpected:\n%s\nGot:\n%s", i, expectedCommands[i], command)
			}
		}
	})
}
