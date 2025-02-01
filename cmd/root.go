package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cristianoliveira/sway-setter/internal/setters"
	"github.com/jessevdk/go-flags"
)

type CliArgs struct {
	Type  string `short:"t" long:"type" description:"Type 'set_{type}' object to be set analog of 'swaymsg -t get_{type}'"`
	Print bool   `short:"p" long:"print" description:"Prints commands that would be executed. Can be used as input to swaymsg"`
}

func Args() *CliArgs {
	var opts CliArgs

	if len(os.Args) <= 1 {
		flags.ParseArgs(&opts, []string{"-h"})
		os.Exit(1)
	}

	args := os.Args[1:]
	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		os.Exit(0)
	}

	return &opts
}

func Execute(opts *CliArgs) error {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		err = scanner.Err()
		return fmt.Errorf("Error: failed to read input\n%s", err)
	}

	if opts.Print {
		setters.ConfigureStdout()
	}

	switch opts.Type {
	case "set_workspaces":
		input := ""
		for scanner.Scan() {
			input += scanner.Text() + "\n"
		}

		if len(input) == 0 {
			return fmt.Errorf("Error: no input provided")
		}

		var workspaces []setters.SwayWorkspace
		err = json.Unmarshal([]byte(input), &workspaces)
		if err != nil {
			return fmt.Errorf("Error: failed to parse input content\nHint: make sure to use a file generated by the `swaymsg -t get_workspaces` output")
		}

		setters.SetWorkspaces(workspaces)

		return nil

	default:
		return fmt.Errorf("Error: type `%s` is not supported\nSupported types\n set_workspaces: load workspaces from `swaymsg` output", opts.Type)
	}
}
