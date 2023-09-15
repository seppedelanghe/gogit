package gogit

import (
	"fmt"
)

func Active(args []string) {
  config := LoadConfig()
  if len(config.Profiles) > 0 {
    fmt.Printf("Active git profile: %s\n", config.Active.Name)
  } else {
    fmt.Println("You do not have an profiles setup yet. Run 'gogit add' to add a profile")
  }
}

func List(args []string) {
  config := LoadConfig()
  if len(config.Profiles) == 0 {
    fmt.Println("You do not have an profiles setup yet. Run 'gogit add' to add a profile")
    return
  }

  var content string = "profiles:\n"
  for _, prof := range config.Profiles {
    if prof.Active {
      content = content + fmt.Sprintf("-> %s\n", prof.Name)
    } else {
      content = content + fmt.Sprintf("-  %s\n", prof.Name)
    }
  }
  fmt.Printf(content)
}

func Add(args []string) {
  if len(args) == 0 {
    fmt.Println("Missing argument 'name' for profile")
    return
  }

  config := LoadConfig()
  name := args[0]

  for _, profile := range config.Profiles {
    if name == profile.Name {
      fmt.Printf("Profile with name '%s' already exists\n", name)
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

  // Custom remote
  if len(args) >= 4 {
    settings.RemoteName = args[3]
  }
  
  profile := Profile{
    Name: name,
    Active: false,
    Settings: settings,
  }

  config.Profiles = append(config.Profiles, profile)
  SaveConfig(&config)

  fmt.Print("Added new profile:\n\n")
  fmt.Println(profile.String())
}

func Set(args []string) {
  if len(args) == 0 {
    fmt.Println("Missing argument 'name' for profile")
    return
  }

  config := LoadConfig()
  name := args[0]

  profile := FindProfile(&config, name)
  if profile == nil {
    fmt.Printf("Profile with name '%s' not found\n", name)
    return
  }

  if config.Active.Name == name {
    fmt.Printf("Profile with name '%s' is already active\n", name)
    return
  }

  if config.Active.Settings.RemoteName != profile.Settings.RemoteName {
    fmt.Printf(
      "Active profile '%s' and profile '%s' have different remotes setup, cannot replace one with another", 
      config.Active.Name, profile.Name,
    )
    return
  }

  // SSH
  fmt.Printf("Setting '%s' as active profile\n", name)
  SetActiveHost(config.Active.Settings.RemoteName, config.Active.Name, profile.Name)

  // Git
  if config.SetGit {
    SetGitUser(&profile.Settings)
  }
  
  // Config
  SetActiveProfile(&config, profile.Name)
  SaveConfig(&config)
}

func Remove(args []string) {
  if len(args) == 0 {
    fmt.Println("Missing argument 'name' for profile")
    return
  }

  config := LoadConfig()
  name := args[0]
  profile := FindProfile(&config, name)
  if profile == nil {
    fmt.Printf("profile with name '%s' not found\n", name)
    return
  }

  if config.Active.Name == profile.Name {
    fmt.Printf("Profile '%s' is currently active, cannot remove\n", name)
    return
  }
  
  // Config
  DeleteProfile(&config, profile)

  fmt.Printf("Profile '%s' removed\n", name)
  
}

func Init(args []string) {
  CreateEmptyConfig()
  fmt.Println("GoGit initialized!")
}
