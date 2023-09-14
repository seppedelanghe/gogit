package gogit

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
)

func getGitConfigPath() string {
  return filepath.Join(os.Getenv("HOME"), ".gitconfig")
}

func replaceGitconfigWithTemp() {
  currentPath := getGitConfigPath()
  os.Rename(currentPath, filepath.Join(os.Getenv("HOME"), ".gitconfig.bak"))

  newFile := "./gitconfig.tmp"

  err := os.Rename(newFile, currentPath)
  if err != nil {
    fmt.Printf("failed to copy new git config to %s", currentPath)
  }
}

func loadGitConfigFile() *ini.File {
  cfg, err := ini.Load(getConfigPath())
  if err != nil {
    fmt.Printf("Failed to read gogit config file: %v", err)
    os.Exit(1)
  }

  return cfg
}

func GetGitUser() (name string, email string) {
  f, err := os.Open(getGitConfigPath())
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()
  scanner := bufio.NewScanner(f)

  var reachedUserSection bool
  for scanner.Scan() {
    text := scanner.Text()
    if text == "[user]" {
      reachedUserSection = true
    } else if reachedUserSection && name == email {
      name = strings.TrimSpace(strings.Split(text, "=")[1])
    } else if reachedUserSection {
      email = strings.TrimSpace(strings.Split(text, "=")[1])
      break
    } 
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return
}

func SetGitUser(settings *ProfileSettings) {
  f, err := os.Open(getGitConfigPath())
  if err != nil {
    log.Fatal(err)
  }

  fw, err := os.Create("gitconfig.tmp")
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()
  defer fw.Close()

  scanner := bufio.NewScanner(f)

  var counter int

  for scanner.Scan() {
    text := scanner.Text()
    if text == "[user]" {
      counter = 1
      fw.WriteString(text + "\n")
    } else if counter == 1 {
      fw.WriteString(fmt.Sprintf("\tname = %s\n", settings.GitName))
      counter = 2
    } else if counter == 2 {
      fw.WriteString(fmt.Sprintf("\temail = %s\n", settings.GitEmail))
      counter = 3
    } else {
      fw.WriteString(text + "\n")
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  replaceGitconfigWithTemp()

}
