package setters

import "github.com/cristianoliveira/sway-setter/internal/sway"

func ConfigStdoutConnector() {
	sway.SwayIPCConnector = &sway.StdOutputConnector{}
}

func ConnectToSway() (*sway.SwayMsgConnection, error) {
	swaymsg, err := sway.SwayIPCConnector.Connect()
	if err != nil {
		return nil, err
	}

	return swaymsg, nil
}

type SwayRoot struct {
	Id    int
	Name  string
	Nodes []SwayOutput
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
	Active     bool            `json:"active"`
	Dpms       bool            `json:"dpms"`
	Transform  string          `json:"transform"`
	Rect       *OutputRect     `json:"rect"`
	CurentMode *Mode           `json:"current_mode"`
	Nodes      []SwayWorkspace `json:"nodes"`
}

type SwayWorkspace struct {
	Id      int    `json:"id"`
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
	Name             string                         `json:"name"`
	Focused          bool                           `json:"focused"`
	AppId            string                        `json:"app_id"`
	WindowProperties *SwayContainerWindowProperties `json:"window_properties"`
	Nodes            []SwayContainer                `json:"nodes"`
	Marks            []string                       `json:"marks"`
}
