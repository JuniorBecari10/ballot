package creator

import (
  "fmt"
  
  "ballot/util"
)

var editing *Ballot

type Ballot struct {
  Name string
  Sections []Section  
}

type Section struct {
  Candidates []Candidate
  NumberLength int
}

type Candidate struct {
  Name string
  Vice string
  
  Number string // yes, a string. Because you can set "05" as a number.
}

func NewBallot(name string) *Ballot {
  return &Ballot { Name: name }
}

func CreateMenu() {
  fmt.Print("Enter the ballot name: ")
  util.Scanner.Scan()
  
  editing = NewBallot(util.Scanner.Text())
}

func MainMenu() {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot %s\n", editing.Name)
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Edit sections")
    fmt.Println("2 - Edit ballot name")
    
    fmt.Println("0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        
        break
      
      case "2":
        EditBallotName()
        break
      
      case "0":
        return
    }
  }
}

func EditBallotName() {
  fmt.Print("Enter the ballot name: ")
  util.Scanner.Scan()
  
  editing.Name = util.Scanner.Text()
}