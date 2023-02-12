package creator

import (
  "fmt"
  
  "ballot/util"
)

var editing *Ballot

// test
type Candidate int
type Section int

type Ballot struct {
  name string
  
  candidates []Candidate
  sections []Section
  
  saved bool
}

func NewBallot(name string) *Ballot {
  return &Ballot { name: name, saved: false }
}

func CreateMenu() {
  fmt.Print("Enter the ballot name: ")
  util.Scanner.Scan()
  
  editing = NewBallot(util.Scanner.Text())
  
  fmt.Println(editing)
}

func MainMenu() {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot %s\n", editing.name)
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Edit section")
    fmt.Println("2 - ...")
    
    fmt.Println("0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        break
      
      case "2":
        break
      
      case "0":
        break
    }
  }
}