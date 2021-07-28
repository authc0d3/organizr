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
  params := utils.OrganizrParams{}
  params.SrcPath = flag.String("src", "", "Source directory")
  params.DestPath = flag.String("dest", "", "Output directory")
  params.Recursive = flag.Bool("r", false, "Recursive mode")
  params.CopyMode = flag.Bool("c", false, "Copy files instead of moving them")
  params.PreserveDuplicates = flag.Bool("p", false, "Preserve duplicate files")
  params.ShowAbout = flag.Bool("about", false, "Show about info")
  flag.Parse()

  if *params.DestPath == "" {
    params.DestPath = params.SrcPath
  }

  if *params.SrcPath == "" || *params.ShowAbout {
    PrintAbout()
    os.Exit(0)
  }

  // Let the magic begins ;)
  utils.OrganizeFiles(&config, &params)
}