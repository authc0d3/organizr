package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/udhos/equalfile"
)

var config *Config
var params *OrganizrParams

func createFolder(path string, mode fs.FileMode) bool {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    err := os.Mkdir(path, mode)
    if err != nil {
      fmt.Println("Error creating directory: " + path)
      return false
    }
  }
  return true
}

func copyFile(file string, folder string) bool {
  srcFile, err := os.Open(file)
  if err != nil {
    fmt.Println("Error opening source file " + file)
    return false
  }

  destPath := folder + "/" + GetSubfolder(file, config)
  createFolder(destPath, 0700)
  destFilePath := GetFinalPath(file, destPath + "/" + filepath.Base(file), *params.PreserveDuplicates, 1)

  destFile, err := os.Create(destFilePath)
  if err != nil {
    fmt.Println("Error creating file " + destFilePath)
    return false
  }

  _, err = io.Copy(destFile, srcFile)
  if err != nil {
    fmt.Println("Error copying file " + file + " to " + destPath)
    return false
  }

  err = destFile.Sync()
  if err != nil {
    fmt.Println("Sync error on copied file " + destFilePath)
    return false
  }

  return true
}

func moveFile(file string, folder string) bool {
  destPath := folder + "/" + GetSubfolder(file, config)
  createFolder(destPath, 0700)

  destFilePath := GetFinalPath(file, destPath + "/" + filepath.Base(file), *params.PreserveDuplicates, 1)
  err := os.Rename(file, destFilePath)
  if err != nil {
    fmt.Println("Error moving file " + file + " to " + destPath)
    return false
  }
  return true
}

func equalFiles(pathA string, pathB string) (bool, error) {
  comparator := equalfile.NewMultiple(nil, equalfile.Options{}, sha256.New(), true)
  equal, err := comparator.CompareFile(pathA, pathB)
  return equal, err
}

// Read the source path (recursively or not) and move files
func OrganizeFiles(c *Config, p *OrganizrParams) {
  config = c
  params = p
  err := filepath.Walk(*params.SrcPath, func(filePath string, f os.FileInfo, err error) error {
    if !f.IsDir() {
      if *params.CopyMode {
        copyFile(filePath, *params.DestPath)
      } else {
        moveFile(filePath, *params.DestPath)
      }
    } else {
      // Skip folder if is not source and recursive mode is disabled
      if !*params.Recursive && filePath != *params.SrcPath {
        return filepath.SkipDir
      }
    }
    return nil
  })

  if err != nil {
    panic("Error reading path: " + *params.SrcPath)
  }
}