package cmd

import (
  "fmt"
  "os"
)

func PrintHelp() {
  fmt.Println("")
  fmt.Println("Organizr " + os.Getenv("version") + ", by Authc0d3")
  fmt.Println("")
  fmt.Println("How to use:")
  fmt.Println("  organizr -src=<source directory> [-dist=<output directory>]")
  fmt.Println("")
  fmt.Println("Flags:")
  fmt.Println("  -h   Show help")
  fmt.Println("  -r   Recursive mode")
}