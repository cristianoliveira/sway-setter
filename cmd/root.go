package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func ScanStdin() (string, error) {
	input := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	if scanner.Err() != nil {
		err := scanner.Err()
		return "", fmt.Errorf("Error: failed to read input\n%s", err)
	}

	if len(input) == 0 {
		return "", fmt.Errorf("Error: no input provided")
	}

	return input, nil
}

func SharedFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().BoolP("print", "p", false, "Prints commands that would be executed. Can be used as input to swaymsg")

	return cmd
}

func ExecuteCommand() error {
	var err error
	rootCmd := &cobra.Command{
		Use:   "sway-setter",
		Short: "The missing cli to use swaymsg 'getters' outputs",
	}

	rootCmd.AddCommand(SharedFlags(OutputCmd()))
	rootCmd.AddCommand(SharedFlags(WorkspacesCmd()))
	rootCmd.AddCommand(SharedFlags(ContainersCmd()))

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return err
}
