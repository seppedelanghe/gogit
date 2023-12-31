package main

import (
	"fmt"
	"os"

	"github.com/seppedelanghe/gogit/gogit"
)


func main() {
  funcs := map[string]func(args []string) {
    "active": gogit.Active,
    "add": gogit.Add,
    "set": gogit.Set,
    "remove": gogit.Remove,
    "init": gogit.Init,
    "list": gogit.List,
  }

  if len(os.Args) == 1 {
    fmt.Println("missing action")
    os.Exit(1)
  }

  actionName := os.Args[1]
  action, exists := funcs[actionName]
  if !exists {
    fmt.Printf("gogit action '%s' does not exists\n", actionName)
    os.Exit(1)
  }

  args := os.Args[2:]
  action(args)
}
