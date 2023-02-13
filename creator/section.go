package creator 

import (
  "fmt"
  
  "ballot/util"
  "strconv"
)

type Section struct {
  Name string
  Candidates []*Candidate
  NumberLength int
}

func NewSection(name string, numberLen int) *Section {
  return &Section { Name: name, NumberLength: numberLen }
}

type Candidate struct {
  Name string
  Vice string
  
  Number string // yes, a string. Because you can set "05" as a number.
}

func NewCandidate(name string, vice string, number string) *Candidate {
  return &Candidate { Name: name, Vice: vice, Number: number }
}

func EditSection(section *Section) {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot %s / section %s\n", editing.Name, section.Name)
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Edit section name")
    fmt.Println("2 - Edit candidates")
    
    fmt.Println("0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        EditSectionName(section)
        break
      
      case "2":
        
        break
      
      case "0":
        return
    }
  }
}

func EditSectionName(s *Section) {
  fmt.Print("Enter the section name: ")
  util.Scanner.Scan()
  
  s.Name = util.Scanner.Text()
}

func EditCandidates(s *Section) {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot %s / section %s / candidates\n", s.Name)
    
    fmt.Println("\nCandidates:\n")
    
    if len(s.Candidates) > 0 {
      for _, c := range s.Candidates {
       fmt.Printf("%s | vice: %s | number: %s\n", c.Name, c.Vice, c.Number); 
      }
    } else {
      fmt.Println("There are no candidates.")
    }
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Add new candidate")
    
    if len(s.Candidates) > 0 {
      fmt.Println("2 - Edit an existing candidate")
    }
    
    fmt.Println("0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        AddCandidate(s)
        break
      
      case "2":
        if len(editing.Sections) <= 0 {
          break
        }
        
        name := GetCandidateName()
        var c *Candidate = nil
        
        for _, cd := range s.Candidates {
          if cd.Name == name {
            c = cd
          }
        }
        
        if c == nil {
          break
        }
        
        //EditCandidate(s)
        break
      
      case "0":
        return
    }
  }
}

// ----

func AddCandidate(s *Section) {
  fmt.Print("Enter the candidate name: ")
  util.Scanner.Scan()
  name := util.Scanner.Text()
  
  fmt.Print("Enter the vice: ")
  util.Scanner.Scan()
  vice := util.Scanner.Text()
  
  fmt.Print("Enter the number: ")
  util.Scanner.Scan()
  number := util.Scanner.Text()
  
  _, err := strconv.Atoi(number)
  
  if err != nil {
    return
  }
  
  for _, c := range s.Candidates {
    if c.Number == number {
      return
    }
  }
  
  s.Candidates = append(s.Candidates, NewCandidate(name, vice, number))
}

func GetCandidateName() string {
  fmt.Print("Enter the candidate name: ")
  util.Scanner.Scan()
  
  return util.Scanner.Text()
}