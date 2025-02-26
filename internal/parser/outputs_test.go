package parser

import (
	"testing"
)

func TestOutputSetter(t *testing.T) {
	t.Run("SetOutputPosition", func(t *testing.T) {
		swayOutputs := []SwayOutput{
			{
				Name: "eDP-1",
				Rect: &Rect{
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

				Rect: &Rect{
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

				Rect: &Rect{
					X:      1920,
					Y:      0,
					Width:  1280,
					Height: 1024,
				},
			},
			{
				Name: "HDMI-A-90",

				Rect: &Rect{
					X:      1920,
					Y:      0,
					Width:  1280,
					Height: 1024,
				},

				Transform: "90",
			},
		}

		expectedCommands := []string{
			"output eDP-1 position 0 0 resolution 1920x1080@60Hz",
			"output HDMI-A-1 position 1920 0 resolution 1280x1024@50Hz",
			"output HDMI-A-2 position 1920 0",
			"output HDMI-A-90 position 1920 0 transform 90",
		}

		commands, err := SetOutputsCommand(swayOutputs)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if len(*commands) != len(expectedCommands) {
			t.Errorf(
				"Expected %d commands to be executed, got %d",
				len(expectedCommands),
				len(*commands),
			)
		}

		for i, command := range *commands {
			if command != expectedCommands[i] {
				t.Errorf("Command:%d\nExpected:\n%s\nGot:\n%s", i, expectedCommands[i], command)
			}
		}
	})
}
