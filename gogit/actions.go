package gogit

import (
	"fmt"
)

func Active(args []string) {
  config := LoadConfig()
  if len(config.Profiles) > 0 {
    fmt.Printf("The current active git profile is: %s\n", config.Active.Name)
  } else {
    fmt.Println("You do not have an profiles setup yet. Run 'gogit add' to add a profile")
  }
}

func Add(args []string) {
  if len(args) == 0 {
    fmt.Println("missing name for profile")
    return
  }

  config := LoadConfig()
  name := args[0]

  for _, profile := range config.Profiles {
    if name == profile.Name {
      fmt.Printf("profile with name '%s' already exists\n", name)
      return
    }
  }

  var settings ProfileSettings
  settings.Fill()
  
  // Custom email
  if len(args) >= 2 {
    settings.GitEmail = args[1]
  }

  // Custom name
  if len(args) >= 3 {
    settings.GitName = args[2]
  }
  
  profile := Profile{
    Name: name,
    Active: false,
    Settings: settings,
  }

  config.Profiles = append(config.Profiles, profile)
  SaveConfig(config)

  fmt.Println("added new profile:\n ")
  fmt.Println(profile.String())
}

func Set(args []string) {
  if len(args) == 0 {
    fmt.Println("missing name for profile")
    return
  }

  config := LoadConfig()
  name := args[0]

  profile := FindProfile(&config, name)
  if profile == nil {
    fmt.Printf("profile with name '%s' not found\n", name)
    return
  }

  if config.Active.Name == name {
    fmt.Printf("profile with name '%s' is already active\n", name)
    return
  }
  
  fmt.Printf("setting '%s' as active profile\n", name)
  SetActiveHost(config.Active.host(), profile.host())
  
  SetActiveProfile(&config, profile.Name)
  SaveConfig(config)
}

func Drop(args []string) {
  fmt.Println("drop not implemented yet")
}

func Init(args []string) {
  CreateEmptyConfig()
}
