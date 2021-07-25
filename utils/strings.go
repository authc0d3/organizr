package utils

import (
  "crypto/sha256"
  "os"
  "path/filepath"
  "strconv"
  "strings"
  "github.com/udhos/equalfile"
)

// Get folder name by file type
// TODO: Refactor, change switch by map and regex (best approach)
func GetSubfolder(file string) string {
  ext := strings.Replace(strings.ToLower(filepath.Ext(file)), ".", "", -1)
  switch ext {
    case "doc", "docx", "xls", "xlsx", "ppt", "pptx", "pdf", "txt", "odt", "ods", "odp", "odg":
      return "Documents"
    case "jpg", "jpeg", "png", "gif", "bmp", "tiff":
      return "Images"
    case "mp3", "ogg", "wma", "wav":
      return "Audios"
    case "mp4", "mkv", "avi", "mov", "mpeg", "wmv":
      return "Videos"
    case "exe", "msi", "so", "apk", "ipa":
      return "Applications"
    default:
      return strings.ToUpper(ext)
  }
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