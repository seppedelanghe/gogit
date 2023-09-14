package gogit

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-ini/ini"
)


func getConfigPath() string {
  if os.Getenv("GOGIT_ENV") == "develop" {
    return "./gogit.ini"
  }
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

  // Preferences
  config.SetGit = cfg.Section("preferences").Key("set-git-user").MustBool(false)

  // Profiles
  for _, name := range cfg.Section("profiles").KeyStrings() {
    val, err := cfg.Section("profiles").Key(name).Bool()
    if err != nil {
      val = false
    }

    cfgProfSection := cfg.Section(fmt.Sprintf("profile.%s", name))

    account := Profile{
      Name: name,
      Active: val,
      Settings: ProfileSettings{
        RemoteName: cfgProfSection.Key("remote").String(),
        GitName: cfgProfSection.Key("username").String(),
        GitEmail: cfgProfSection.Key("email").String(),
      },
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

  // Preferences
  cfg.Section("preferences").Key("set-git-user").SetValue(strconv.FormatBool(config.SetGit))

  // Profiles
  for _, profile := range config.Profiles {
    str_bool := strconv.FormatBool(profile.Active)
    cfg.Section("profiles").Key(profile.Name).SetValue(str_bool)
    
    sectionKey := fmt.Sprintf("profile.%s", profile.Name)
    cfg.Section(sectionKey).Key("remote").SetValue(profile.Settings.RemoteName)
    cfg.Section(sectionKey).Key("username").SetValue(profile.Settings.GitName)
    cfg.Section(sectionKey).Key("email").SetValue(profile.Settings.GitEmail)
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
  cfg.Section("preferences").Key("set-git-user").SetValue("false")

  cfg.SaveTo(getConfigPath())
}
