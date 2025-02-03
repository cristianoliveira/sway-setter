package parser

import (
	"testing"

	"github.com/cristianoliveira/sway-setter/internal/sway"
	"github.com/cristianoliveira/sway-setter/internal/testutils"
)

func TestContainerSetter(t *testing.T) {
	t.Run("SetContainers", func(t *testing.T) {
		con := testutils.MockedConnector{}
		sway.SwayIPCConnector = &con

		swayWorkspaces := []SwayWorkspace{
			{
				Name: "1",
				Nodes: []SwayContainer{
					{
						AppId: "1",
					},
					{
						AppId: "",
						WindowProperties: &SwayContainerWindowProperties{
							Title: "foobar",
						},
					},
					{
						AppId: "",
						Name:  "Window with special chars",
						WindowProperties: &SwayContainerWindowProperties{
							Title: "(foo) - foobar \\ bar",
						},
					},
				},
			},
			{
				Name: "2",
				Nodes: []SwayContainer{
					{
						AppId: "",
						WindowProperties: &SwayContainerWindowProperties{
							Class: "foobarclass",
						},
					},
					{
						AppId: "",
						Marks: []string{"setter:1"},
					},
				},
				FloatinNodes: []SwayContainer{
					{
						AppId: "fn1",
						Rect: &Rect{
							X: 0,
							Y: 200,
						},
					},
				},
			},
		}

		expectedCommands := []string{
			"[app_id=\"1\"] floating disable; [app_id=\"1\"] move container to workspace 1",
			"[title=\"foobar\"] floating disable; [title=\"foobar\"] move container to workspace 1",
			"[title=\"\\(foo\\) - foobar \\\\ bar\"] floating disable; [title=\"\\(foo\\) - foobar \\\\ bar\"] move container to workspace 1",
			"[class=\"foobarclass\"] floating disable; [class=\"foobarclass\"] move container to workspace 2",
			"[con_mark=\"setter:1\"] floating disable; [con_mark=\"setter:1\"] move container to workspace 2",
			"[app_id=\"fn1\"] floating enable; [app_id=\"fn1\"] move container to workspace 2",
			"[app_id=\"fn1\"] move absolute position 0 200",
		}

		commands, err := SetContainersCommand(swayWorkspaces)
		if err != nil {
			t.Errorf("\nExpected no error, got: %s", err)
		}

		if len(*commands) != len(expectedCommands) {
			t.Fatalf(
				"Expected %d commands to be executed, got %d",
				len(expectedCommands),
				len(*commands),
			)
		}

		for i, command := range *commands {
			if command != expectedCommands[i] {
				t.Errorf("\nExpected: %s\nReceived: %s", expectedCommands[i], command)
			}
		}
	})
}
