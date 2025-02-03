package setters

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
			},
		}

		expectedCommands := []string{
			"[app_id=\"1\"] move container to workspace 1",
			"[title=\"foobar\"] move container to workspace 1",
			"[class=\"foobarclass\"] move container to workspace 2",
			"[con_mark=\"setter:1\"] move container to workspace 2",
		}

		err := SetContainers(swayWorkspaces)
		if err != nil {
			t.Errorf("Expected no error, got: %s", err)
		}

		if len(con.CommandsHistory) != len(expectedCommands) {
			t.Fatalf(
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
