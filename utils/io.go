package utils

import (
  "fmt"
  "io"
  "io/fs"
  "os"
  "path/filepath"
)

func createFolder(path string, mode fs.FileMode) bool {
  // destPath := folder + "/" + GetSubfolder(file)
  if _, err := os.Stat(path); os.IsNotExist(err) {
    err := os.Mkdir(path, mode)
    if err != nil {
      fmt.Println("Error creating directory: " + path)
      return false
    }
  }
  return true
}

// Copy file
func copyFile(file string, folder string) bool {
  srcFile, err := os.Open(file)
  if err != nil {
    fmt.Println("Error opening source file " + file)
    return false
  }

  destPath := folder + "/" + GetSubfolder(file)
  createFolder(destPath, 0700)
  destFilePath := GetFinalPath(file, destPath + "/" + filepath.Base(file), 1)

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

// Move file to another folder
func moveFile(file string, folder string) bool {
  destPath := folder + "/" + GetSubfolder(file)
  createFolder(destPath, 0700)

  destFilePath := GetFinalPath(file, destPath + "/" + filepath.Base(file), 1)
  err := os.Rename(file, destFilePath)
  if err != nil {
    fmt.Println("Error moving file " + file + " to " + destPath)
    return false
  }
  return true
}

// Read the source path (recursively or not) and move files
func OrganizeFiles(srcPath string, destPath string, recursive bool, copyMode bool) {
  err := filepath.Walk(srcPath, func(filePath string, f os.FileInfo, err error) error {
    if !f.IsDir() {
      if copyMode {
        copyFile(filePath, destPath)
      } else {
        moveFile(filePath, destPath)
      }
    } else {
      // Skip folder if is not source and recursive mode is disabled
      if !recursive && filePath != srcPath {
        return filepath.SkipDir
      }
    }
    return nil
  })

  if err != nil {
    panic("Error reading path: " + srcPath)
  }
}