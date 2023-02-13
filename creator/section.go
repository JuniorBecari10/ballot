package creator 

import (
  "fmt"
  
  "ballot/util"
  "strings"
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
        EditCandidates(section)
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
    fmt.Printf("Editing ballot %s / section %s / candidates\n", editing.Name, s.Name)
    
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
        
        number := GetCandidateNumber()
        var c *Candidate = nil
        
        for _, cd := range s.Candidates {
          if strings.ToLower(cd.Number) == strings.ToLower(number) {
            c = cd
          }
        }
        
        if c == nil {
          break
        }
        
        EditCandidate(s, c)
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

func GetCandidateNumber() string {
  fmt.Print("Enter the candidate number: ")
  util.Scanner.Scan()
  
  return util.Scanner.Text()
}