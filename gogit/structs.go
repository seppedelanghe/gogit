package gogit

import (
  "fmt"
)

type ProfileSettings struct {
  RemoteName string
  GitName string
  GitEmail string
}

func (settings *ProfileSettings) Fill() {
  settings.RemoteName = "github.com"
  
  gitname, gitemail := GetGitUser()
  settings.GitName = gitname
  settings.GitEmail = gitemail
}

type Profile struct {
  Name string
  Active bool
  Settings ProfileSettings
}

func (prof *Profile) host() string {
  return fmt.Sprintf("github.com-%s", prof.Name)
}

type Config struct {
  Profiles []Profile
  Active Profile
}
