package gogit

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

  

func getSshConfigPath() (string) {
  return filepath.Join(os.Getenv("HOME"), ".ssh", "config")
}

func MoveTempFile() {
  currentPath := getSshConfigPath()
  newPath := filepath.Join(os.Getenv("HOME"), ".ssh", "config.bak")

  err := os.Rename(currentPath, newPath)
  if err != nil {
    fmt.Println("Failed to create backup of SSH config file, aborting...")
    return
  }

  err = os.Rename("config.tmp", currentPath)
  if err != nil {
    fmt.Println("Failed to move new SSH config file, to .ssh directory. Manual movement is recommended")
    return
  }

  fmt.Println("New ssh config enabled")

}

func SetActiveHost(activehost string, newhost string) {
  f, err := os.Open(getSshConfigPath())
  if err != nil {
    log.Fatal(err)
  }

  fnew, err := os.Create("config.tmp")
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()
  defer fnew.Close()

  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    text := scanner.Text()
    if text == fmt.Sprintf("Host %s", newhost) {
      fnew.WriteString("Host github.com\n")
    } else if text == "Host github.com" {
      fnew.WriteString(fmt.Sprintf("Host %s\n", activehost))
    } else {
      fnew.WriteString(text + "\n")
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  MoveTempFile()
}


