package cmd

import (
  "flag"
  "os"
  "github.com/authc0d3/organizr/utils"
)

// Globals
var srcPath *string
var destPath *string

func Organizr() {
  // Get and validate flags
  srcPath = flag.String("src", "", "Source directory")
  destPath = flag.String("dest", "", "Output directory")
  recursive := flag.Bool("r", false, "Recursive mode")
  help := flag.Bool("h", false, "Show help")
  flag.Parse()

  if *destPath == "" {
    destPath = srcPath
  }

  if *srcPath == "" || *help {
    PrintHelp()
    os.Exit(0)
  }

  // Let the magic begins ;)
  utils.OrganizeFiles(*srcPath, *destPath, *recursive)
}