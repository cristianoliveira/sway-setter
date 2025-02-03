package main

import (
	"os"

	"github.com/cristianoliveira/sway-setter/cmd"
)

func main() {
	err := cmd.ExecuteCommand()
	if err != nil {
		os.Exit(1)
	}
}
