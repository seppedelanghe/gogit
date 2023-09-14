package gogit

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-ini/ini"
)

type Profile struct {
  Name string
  Active bool
}

func (prof *Profile) host() string {
  return fmt.Sprintf("github.com-%s", prof.Name)
}

type Config struct {
  Profiles []Profile
  Active Profile
}

func getConfigPath() string {
  return filepath.Join(os.Getenv("HOME"), "gogit.ini")
}

func loadConfigFile() (*ini.File) {
  cfg, err := ini.Load(getConfigPath())
  if err != nil {
    fmt.Printf("Failed to read gogit config file: %v", err)
    os.Exit(1)
  }
  
  return cfg
}

func LoadConfig() (config Config) {
  cfg := loadConfigFile()

  for _, name := range cfg.Section("profiles").KeyStrings() {
    val, err := cfg.Section("profiles").Key(name).Bool()
    if err != nil {
      val = false
    }
    account := Profile{
      Name: name,
      Active: val,
    }
    config.Profiles = append(config.Profiles, account)
    if val {
      config.Active = account
    }
  }

  return
}

func SaveConfig(config Config) {
  cfg := loadConfigFile()

  for _, profile := range config.Profiles {
    str_bool := strconv.FormatBool(profile.Active)
    cfg.Section("profiles").Key(profile.Name).SetValue(str_bool)
  }

  cfg.SaveTo(getConfigPath())
}


func SetActiveProfile(config *Config, name string) {
  for i, prof := range config.Profiles {
    if name == prof.Name {
      config.Profiles[i].Active = true
      config.Active = prof
    } else {
      config.Profiles[i].Active = false
    }
  }
}

func FindProfile(config *Config, name string) (profile *Profile) {
  for _, prof := range config.Profiles {
    if name == prof.Name {
      return &prof
    }
  }
  
  return nil
}

func CreateEmptyConfig() {
  cfg := ini.Empty()
  cfg.Section("profiles").SetBody("")
  cfg.SaveTo(getConfigPath())

  fmt.Println("gogit initialized!")
}
