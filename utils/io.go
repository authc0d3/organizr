package utils

import (
  "fmt"
  "os"
  "path/filepath"
)

// Move file to another folder
func moveFile(file string, folder string) {
  destPath := folder + "/" + GetSubfolder(file)
  if _, err := os.Stat(destPath); os.IsNotExist(err) {
    err := os.Mkdir(destPath, 0700)
    if err != nil {
      fmt.Println("Error creating directory: " + destPath)
    }
  }

  destFilePath := GetFinalPath(file, destPath + "/" + filepath.Base(file), 1)
  err := os.Rename(file, destFilePath)
  if err != nil {
    fmt.Println("Error moving file: " + file)
  }
}

// Read the source path (recursively or not) and move files
func OrganizeFiles(srcPath string, destPath string, recursive bool) {
  err := filepath.Walk(srcPath, func(filePath string, f os.FileInfo, err error) error {
    if (!f.IsDir()) {
      moveFile(filePath, destPath);
    } else {
      // Skip folder if is not source and recursive mode is disabled
      if (!recursive && filePath != srcPath) {
        return filepath.SkipDir
      }
    }
    return nil
  })

  if err != nil {
    panic("Error reading path: " + srcPath)
  }
}