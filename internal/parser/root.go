package parser

import (
	"fmt"
)

type SwayRoot struct {
	Id    int          `json:"id"`
	Type  string       `json:"type"`
	Name  string       `json:"name"`
	Nodes []SwayOutput `json:"nodes"`
}

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
	Id         int             `json:"id"`
	Name       string          `json:"name"`
	Type       string          `json:"type"`
	Active     bool            `json:"active"`
	Dpms       bool            `json:"dpms"`
	Transform  string          `json:"transform"`
	Rect       *OutputRect     `json:"rect"`
	CurentMode *Mode           `json:"current_mode"`
	Nodes      []SwayWorkspace `json:"nodes"`
}

type SwayWorkspace struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Output  string `json:"output"`
	Focused bool   `json:"focused"`
	Nodes   []SwayContainer
}

type SwayContainerWindowProperties struct {
	Title    string `json:"title"`
	Class    string `json:"class"`
	Instance string `json:"instance"`
}

type SwayContainer struct {
	Id               int                            `json:"id"`
	Type             string                         `json:"type"`
	Name             string                         `json:"name"`
	Focused          bool                           `json:"focused"`
	AppId            string                         `json:"app_id"`
	WindowProperties *SwayContainerWindowProperties `json:"window_properties"`
	Nodes            []SwayContainer                `json:"nodes"`
	Marks            []string                       `json:"marks"`
}

func CollectWorkspaces(node SwayRoot) ([]SwayWorkspace, error) {
	if node.Type != "root" {
		return nil, fmt.Errorf("the node is not a root node")
	}

	if len(node.Nodes) == 0 {
		return nil, fmt.Errorf("the root doesn't have any outputs")
	}

	workspaces := []SwayWorkspace{}
	for i, outputs := range node.Nodes {
		if outputs.Type != "output" {
			return nil, fmt.Errorf("output %d is not an output node", i)
		}

		for ii, workspace := range outputs.Nodes {
			if workspace.Type != "workspace" {
				return nil, fmt.Errorf("workspace %d in output %d is not a workspace node", ii, i)
			}

			workspaces = append(workspaces, workspace)
		}
	}

	return workspaces, nil
}
