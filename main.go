package main

import (
  "fmt"
  "os"
  
  "ballot/creator"
  "ballot/util"
)

func main() {
  MainMenu()
}

func MainMenu() {
  for {
    util.Clear()
    util.PrintName()
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Create new ballot box")
    fmt.Println("2 - Open an existing ballot box")
    
    fmt.Println("0 - Exit")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        creator.CreateMenu()
        creator.MainMenu()
        break
      
      case "2":
        break
      
      case "0":
        util.Clear()
        os.Exit(0)
    }
  }
}