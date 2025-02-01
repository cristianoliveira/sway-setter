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
		Type   string `short:"t" long:"type" description:"Type 'set_{type}' object to be set analog of 'swaymsg -t get_{type}'"`
		DryRun bool   `long:"dry-run" description:"Dry run mode"`
	}

	args := os.Args[1:]
	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		if len(os.Args) <= 1 {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		fmt.Println("Error: failed to read input")
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.DryRun {
		setters.ConfigDryRun()
	}

	if len(opts.Type) > 0 {
		if opts.Type == "set_workspaces" {
			input := ""
			for scanner.Scan() {
				input += scanner.Text() + "\n"
			}

			if len(input) == 0 {
				fmt.Println("No input provided")
				os.Exit(1)
			}

			var workspaces []setters.SwayWorkspace
			json.Unmarshal([]byte(input), &workspaces)
			setters.SetWorkspaces(workspaces)
			return
		}

		fmt.Printf("Error: type `%s` is not supported\n", opts.Type)
		fmt.Println("Supported types")
		fmt.Println(" set_workspaces: load workspaces from `swaymsg` output")
		os.Exit(1)
	}

	flags.ParseArgs(&opts, []string{"-h"})
	os.Exit(0)
}
