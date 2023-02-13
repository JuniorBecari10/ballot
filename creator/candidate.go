package creator

import (
  "fmt"
  
  "ballot/util"
  "strconv"
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
    
    fmt.Println("\nChoose a field to edit:\n")
    
    fmt.Printf("1 - Name: %s\n", c.Name)
    fmt.Printf("2 - Vice: %s\n", c.Vice)
    fmt.Printf("3 - Number: %s\n", c.Number)
    
    fmt.Println("0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        fmt.Print("Enter the new candidate name: ")
        util.Scanner.Scan()
        
        name := util.Scanner.Text()
        c.Name = name
        
        break
      
      case "2":
        fmt.Print("Enter the new candidate vice: ")
        util.Scanner.Scan()
        
        vice := util.Scanner.Text()
        c.Vice = vice
        break
      
      case "3":
        fmt.Print("Enter the new candidate number: ")
        util.Scanner.Scan()
        
        number := util.Scanner.Text()
        
        _, err := strconv.Atoi(number)
  
        if err != nil {
          break
        }
        
        for _, c := range s.Candidates {
          if c.Number == number {
            break
          }
        }
        
        if len(number) != s.NumberLength {
          break
        }
        
        c.Number = number
        break
      
      case "0":
        return
    }
  }
}