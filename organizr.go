package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/udhos/equalfile"
)

// Constants
const version = "0.23"

// Globals
var srcPath *string
var destPath *string

// Print help
func printHelp() {
  fmt.Println("")
  fmt.Println("Organizr " + version + ", by Authc0d3")
  fmt.Println("")
  fmt.Println("How to use:")
  fmt.Println("  organizr -src=<source directory> [-dist=<output directory>]")
  fmt.Println("")
  fmt.Println("Flags:")
  fmt.Println("  -h   Show help")
  fmt.Println("  -r   Recursive mode")
}

// Get folder name by file type
// TODO: Refactor, change switch by map and regex (best approach)
func getSubfolder(file string) string {
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
func getFinalPath(currentFilePath string, futureFilePath string, iteration int) string {
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
    return getFinalPath(currentFilePath, fileNewName + " (" + strconv.Itoa(iteration) + ")" + ext, iteration + 1)
  }
  return futureFilePath
}

// Move file to another folder
func moveFile(file string, folder string) {
  destPath := folder + "/" + getSubfolder(file)
  if _, err := os.Stat(destPath); os.IsNotExist(err) {
    err := os.Mkdir(destPath, 0700)
    if err != nil {
      fmt.Println("Error creating directory: " + destPath)
    }
  }

  destFilePath := getFinalPath(file, destPath + "/" + filepath.Base(file), 1)
  err := os.Rename(file, destFilePath)
  if err != nil {
    fmt.Println("Error moving file: " + file)
  }
}

// Read the source path (recursively or not) and move files
func organizeFiles(srcPath string, destPath string, recursive bool) {
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

func main() {
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
    printHelp()
    os.Exit(0)
  }

  // Let the magic begins ;)
  organizeFiles(*srcPath, *destPath, *recursive)
}