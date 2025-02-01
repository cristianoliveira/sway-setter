package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cristianoliveira/sway-setter/internal/setters"
	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
    Type string `short:"t" long:"type" description:"Type of input"`
	}

  args := os.Args[1:]

	args, err := flags.ParseArgs(&opts, args)

	if err != nil {
		panic(err)
	}

  scanner := bufio.NewScanner(os.Stdin)
  input := ""
  for scanner.Scan() {
    input += scanner.Text() + "\n"
  }

  if len(input) == 0 {
    fmt.Println("No input provided")
    os.Exit(1)
  }

  fmt.Println(input)

  if scanner.Err() != nil {
    panic(err)
  }

  var workspaces []setters.SwayWorkspace
  json.Unmarshal([]byte(input), &workspaces)
  setters.SetWorkspaces(workspaces)
}
