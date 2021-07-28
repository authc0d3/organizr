package main

import (
	"os"

	"github.com/authc0d3/organizr/cmd"
)

// Constants
const version = "0.26"

func main() {
  os.Setenv("version", version)

  // Let the magic begins ;)
  cmd.Organizr()
}