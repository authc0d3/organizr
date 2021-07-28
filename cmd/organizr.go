package cmd

import (
	"flag"
	"os"

	"github.com/authc0d3/organizr/utils"
)

func Organizr() {
  // Load context
  context := utils.Context{}
  context.Config = utils.GetConfig()
  context.SrcPath = flag.String("src", "", "Source directory")
  context.DestPath = flag.String("dest", "", "Output directory")
  context.Recursive = flag.Bool("r", false, "Recursive mode")
  context.CopyMode = flag.Bool("c", false, "Copy files instead of moving them")
  context.PreserveDuplicates = flag.Bool("p", false, "Preserve duplicate files")
  context.ShowAbout = flag.Bool("about", false, "Show about info")
  flag.Parse()

  if *context.DestPath == "" {
    context.DestPath = context.SrcPath
  }

  if *context.SrcPath == "" || *context.ShowAbout {
    PrintAbout()
    os.Exit(0)
  }

  // Let the magic begins ;)
  utils.OrganizeFiles(&context)
}