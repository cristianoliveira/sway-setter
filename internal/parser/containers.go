package parser

import (
	"fmt"
	"regexp"
	"strings"
)

func escapeChars(str string) string {
	return regexp.QuoteMeta(str)
}

func workspaceToCommand(workspace SwayWorkspace) (string, error) {
	if len(workspace.Name) == 0 {
		return "", fmt.Errorf("Error: workspace name is empty")
	}

	command := fmt.Sprintf("move container to workspace %s", workspace.Name)
	return command, nil
}

// containerToCommand returns a command to send the container to the workspace
// in the given format: "[app_id=foobar]"
// It applies the following precedence:
// 0. mark[setter:<id>] (highest precedence for containers with generic properties)
// 1. app_id
// 2. window_title
// 3. window_class
func containerToCommand(container SwayContainer) (*[]string, error) {
	commands := []string{}
	if len(container.Nodes) > 0 {
		for _, subContainer := range container.Nodes {
			subContainerCmd, err := containerToCommand(subContainer)
			if err != nil {
				return nil, err
			}

			if subContainerCmd != nil {
				commands = append(commands, *subContainerCmd...)
			}
		}

		return &commands, nil
	}

	if len(container.Marks) > 0 {
		for _, mark := range container.Marks {
			if strings.Contains(mark, "setter:") {
				cmd := fmt.Sprintf("[con_mark=\"%s\"]", escapeChars(mark))
				commands = append(commands, cmd)
				return &commands, nil
			}
		}
	}

	if len(container.AppId) > 0 {
		cmd := fmt.Sprintf("[app_id=\"%s\"]", escapeChars(container.AppId))
		commands = append(commands, cmd)
		return &commands, nil
	}

	if container.WindowProperties != nil {
		if len(container.WindowProperties.Title) > 0 {
			cmd := fmt.Sprintf("[title=\"%s\"]", escapeChars(container.WindowProperties.Title))
			commands = append(commands, cmd)
			return &commands, nil
		}

		if len(container.WindowProperties.Class) > 0 {
			cmd := fmt.Sprintf("[class=\"%s\"]", escapeChars(container.WindowProperties.Class))
			commands = append(commands, cmd)
			return &commands, nil
		}
	}

	return nil, fmt.Errorf(
		"Error: container '%s' does not have app_id, title or class",
		container.Name,
	)
}

// SetContainersCommand configures containers to the provided workpace
// containers may be windows and apps.
// Usually a workspace contains one or more containers.
// To check the containers in your workspaces, run:
// ```bash
//
//	swaymsg -t get_tree' \
//	  | jq '[recurse(.nodes[]?, .floating_nodes[]?) | select(.type == "con")]'
//
// ```
func SetContainersCommand(workspaces []SwayWorkspace) (*[]string, error) {
	var commands []string
	if len(workspaces) == 0 {
		return nil, fmt.Errorf("Error: no workspaces provided")
	}

	for _, workspace := range workspaces {
		moveToWorkspaceCmd, err := workspaceToCommand(workspace)
		if err != nil {
			return nil, err
		}

		for _, container := range workspace.Nodes {
			containerCmd, err := containerToCommand(container)
			if err != nil {
				return nil, fmt.Errorf(
					"Error: failed to parse container '%s' in workspace '%s'\nReason: %s",
					container.Name,
					workspace.Name,
					err.Error(),
				)
			}

			for _, cmd := range *containerCmd {
				command := fmt.Sprintf("%s %s", cmd, moveToWorkspaceCmd)
				commands = append(commands, command)
			}
		}
	}

	return &commands, nil
}
