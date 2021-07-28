package cmd

import (
	"flag"
	"os"

	"github.com/authc0d3/organizr/utils"
)

func Organizr() {
  // Load config
  config := utils.GetConfig()

  // Get and validate flags
  srcPath := flag.String("src", "", "Source directory")
  destPath := flag.String("dest", "", "Output directory")
  recursive := flag.Bool("r", false, "Recursive mode")
  copyMode := flag.Bool("c", false, "Copy files instead of moving them")
  showAbout := flag.Bool("about", false, "Show about info")
  flag.Parse()

  if *destPath == "" {
    destPath = srcPath
  }

  if *srcPath == "" || *showAbout {
    PrintAbout()
    os.Exit(0)
  }

  // Let the magic begins ;)
  utils.OrganizeFiles(&config, *srcPath, *destPath, *recursive, *copyMode)
}