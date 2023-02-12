package util

import (
  "fmt"
  "runtime"
  "os/exec"
  "os"
  "bufio"
)

var Scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

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

func PrintName() {
  fmt.Println("Ballot Box Creator")
}