package creator

import (
  "fmt"
  
  "ballot/util"
)

type Candidate struct {
  Name string
  Vice string
  
  Number string // yes, a string. Because you can set "05" as a number.
}

func NewCandidate(name string, vice string, number string) *Candidate {
  return &Candidate { Name: name, Vice: vice, Number: number }
}

func EditCandidate(s *Section, c *Candidate) {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot %s / section %s / candidate %s\n", editing.Name, s.Name, c.Name)
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Edit section name")
    fmt.Println("2 - Edit candidates")
    
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
        return
    }
  }
}