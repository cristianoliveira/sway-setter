package main

import (
	"fmt"
	"os"

	"github.com/cristianoliveira/sway-setter/cmd"
)

func main() {
  opts := cmd.Args()

  err := cmd.Execute(opts)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
