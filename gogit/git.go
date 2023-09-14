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

