package main

import (
  "runtime"
  "os/exec"
  "os"
)

func Clear() {
  switch runtime.GOOS {
    case "linux":
      cmd := exec.Command("clear")
      cmd.Stdout = os.Stdout
      cmd.Run()
      break
    
    case "windows":
      cmd := exec.Command("cmd", "/c", "cls")
      cmd.Stdout = os.Stdout
      cmd.Run()
      break
  }
}