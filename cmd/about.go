package cmd

import (
  "fmt"
  "os"
)

func PrintAbout() {
  fmt.Println("")
  fmt.Println("Organizr " + os.Getenv("version") + ", written in Go by Authc0d3")
  fmt.Println("")
  fmt.Println("How to use:")
  fmt.Println("  organizr -src=<source directory> [-dist=<output directory>]")
  fmt.Println("")
  fmt.Println("Flags:")
  fmt.Println("  -r   Recursive mode")
  fmt.Println("  -c   Copy files instead of moving them")
}