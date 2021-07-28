package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Get folder name by file type acording to context
func GetSubfolder(file string, context *Context) string {
  folders, defaultFolder := context.GetOutputConfig()
  ext := strings.Replace(strings.ToLower(filepath.Ext(file)), ".", "", -1)
  for _, folder := range folders {
    extensions := strings.Split(folder.Ext, ",")
    for _, extension := range extensions {
      if ext == extension {
        return folder.Folder
      }
    }
  }
  if defaultFolder != "" {
    return defaultFolder
  }
  return strings.ToUpper(ext)
}

// Find safe output filename
func GetFinalPath(currentFilePath string, futureFilePath string, preserveDuplicates bool, iteration int) string {
  _, err := os.Stat(futureFilePath)
  if err == nil {
    // If preserveDuplicates is false and output file is same that source file
    // return same name to replace one for another
    if !preserveDuplicates {
      equal, _ := equalFiles(currentFilePath, futureFilePath)
      if equal {
        return futureFilePath;
      }
    }

    // If files are not the same, remove final "(index)" part of file
    // to add it in the next call incremented by 1
    ext := filepath.Ext(futureFilePath)
    fileNewName := futureFilePath[0:len(futureFilePath)-len(ext)]
    fileEnding := " (" + strconv.Itoa(iteration - 1) + ")"
    fileEndingIndex := strings.LastIndex(fileNewName, fileEnding)
    if fileEndingIndex == len(fileNewName) - len(fileEnding) {
      fileNewName = fileNewName[:fileEndingIndex]
    }

    // Call recursively until find a valid name for file
    return GetFinalPath(currentFilePath, fileNewName + " (" + strconv.Itoa(iteration) + ")" + ext, preserveDuplicates, iteration + 1)
  }
  return futureFilePath
}