package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Folder struct {
  Ext    string `json:"ext"`
  Folder string `json:"folder"`
}

type Output struct {
  DefaultFolder string `json:"defaultFolder"`
  Folders []Folder `json:"folders"`
}

type Config struct {
  ExtendsDefault bool `json:"extendsDefault"`
  Output Output `json:"output"`
}

func (c Config) GetOutputConfig() ([]Folder, string) {
  return c.Output.Folders, c.Output.DefaultFolder
}

func GetDefaultConfig() Config {
  defaultConfig := &Config{
    Output: Output {
      Folders: []Folder{},
    },
  }
  defaultConfig.Output.Folders =
    append(defaultConfig.Output.Folders, Folder{ Ext: "doc,docx,xls,xlsx,ppt,pptx,pdf,txt,odt,ods,odp,odg", Folder: "Documents" })
  defaultConfig.Output.Folders =
    append(defaultConfig.Output.Folders, Folder{ Ext: "jpg,jpeg,png,gif,bmp,tiff", Folder: "Images" })
  defaultConfig.Output.Folders =
    append(defaultConfig.Output.Folders, Folder{ Ext: "mp3,ogg,wma,wav", Folder: "Audios" })
  defaultConfig.Output.Folders =
    append(defaultConfig.Output.Folders, Folder{ Ext: "mp4,mkv,avi,mov,mpeg,wmv", Folder: "Videos" })
  defaultConfig.Output.Folders =
    append(defaultConfig.Output.Folders, Folder{ Ext: "exe,msi,so,apk,ipa", Folder: "Applications" })
  return *defaultConfig
}

func GetConfig() Config {
  var config Config

  currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    return GetDefaultConfig()
  }

  configFile, err := os.Open(currentDir + "/config.json")
  if err != nil {
    return GetDefaultConfig()
  }

  bytes, err := ioutil.ReadAll(configFile)
  if err != nil {
    fmt.Println("Error reading config.json file")
    return GetDefaultConfig()
  }

  err = json.Unmarshal(bytes, &config)
  if err != nil {
    fmt.Println("Error unmarshalling config.json file")
    return GetDefaultConfig()
  }

  // Extends custom config with default config
  if config.ExtendsDefault {
    defaults := GetDefaultConfig()

    // By setting the default folders after the custom folders
    // we make sure that the first ones have priority
    config.Output.Folders = append(config.Output.Folders, defaults.Output.Folders...)
  }

  configFile.Close()
  return config
}