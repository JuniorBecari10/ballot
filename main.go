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

    if creator.Editing != nil {
      fmt.Printf("You are currently editing ballot box %s.\n", creator.Editing.Name)
    }
    
    util.PrintErrMsg()
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Create new ballot box")
    fmt.Println("2 - Open an existing ballot box")

    if creator.Editing != nil {
      fmt.Printf("3 - Continue editing ballot box %s\n", creator.Editing.Name)
    }
    
    fmt.Println("0 - Exit")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        if creator.CreateMenu() {
          creator.MainMenu()
        }
        break
      
      case "2":
        filename := creator.LoadBallot()
        ballot, err := util.LoadBallot(filename)
        
        if err != nil {
          util.SetErrMsg(err.Error())
          break
        }
        
        creator.Editing = ballot
        
        break
      
      case "3":
        if creator.Editing != nil {
          creator.MainMenu()
        }
        break

      case "0":
        util.Clear()
        os.Exit(0)
    }
  }
}