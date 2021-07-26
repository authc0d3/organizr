package utils

import (
	"crypto/sha256"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/udhos/equalfile"
)

// Get folder name by file type acording to config
func GetSubfolder(file string) string {
  folders, defaultFolder := GetConfig().GetOutputConfig()
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

// Find safe filename adding a index on the end of the filename to prevent rewrite files
func GetFinalPath(currentFilePath string, futureFilePath string, iteration int) string {
  _, err := os.Stat(futureFilePath)
  if err == nil {
    // If is same file return same name to replace one for another
    comparator := equalfile.NewMultiple(nil, equalfile.Options{}, sha256.New(), true)
    equal, _ := comparator.CompareFile(currentFilePath, futureFilePath)
    if equal {
      return futureFilePath;
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
    return GetFinalPath(currentFilePath, fileNewName + " (" + strconv.Itoa(iteration) + ")" + ext, iteration + 1)
  }
  return futureFilePath
}