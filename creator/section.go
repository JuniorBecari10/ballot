package creator

import (
	"fmt"

	"ballot/util"
	"strconv"
	"strings"
)

func EditSection(section *util.Section) {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot box %s / section %s\n", Editing.Name, section.Name)
    util.PrintErrMsg()
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Edit candidates")
    fmt.Println("2 - Edit section name")
    
    fmt.Println("\n0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        EditCandidates(section)
        break
      
      case "2":
        EditSectionName(section)
        break
      
      case "0":
        return
    }
  }
}

func EditSectionName(s *util.Section) {
  fmt.Print("Enter the new section name: ")
  util.Scanner.Scan()
  
  name := util.Scanner.Text()
  
  if name == "" {
    util.SetErrMsg("The section name cannot be empty!")
    return
  }
  
  for _, s := range Editing.Sections {
    if strings.ToLower(s.Name) == strings.ToLower(name) {
      util.SetErrMsg("There is already a section with this name!")
      return
    }
  }
  
  s.Name = name
}

func EditCandidates(s *util.Section) {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot box %s / section %s / candidates\n", Editing.Name, s.Name)
    util.PrintErrMsg()
    
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
    
    fmt.Println("\n0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        AddCandidate(s)
        break
      
      case "2":
        if len(Editing.Sections) <= 0 {
          break
        }
        
        number := GetCandidateNumber()
        var c *util.Candidate = nil
        
        for _, cd := range s.Candidates {
          if strings.ToLower(cd.Number) == strings.ToLower(number) {
            c = cd
          }
        }
        
        if c == nil {
          util.SetErrMsg("Couldn't find a candidate with the number '" + number + "'.")
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

func AddCandidate(s *util.Section) {
  fmt.Print("Enter the candidate name: ")
  util.Scanner.Scan()
  name := util.Scanner.Text()
  
  if name == "" {
    util.SetErrMsg("The candidate name cannot be empty!")
    return
  }
  
  fmt.Print("Enter the candidate vice: ")
  util.Scanner.Scan()
  vice := util.Scanner.Text()
  
  fmt.Print("Enter the candidate number: ")
  util.Scanner.Scan()
  number := util.Scanner.Text()
  
  _, err := strconv.Atoi(number)
  
  if err != nil {
    util.SetErrMsg("Couldn't process candidate number: '" + number + "'.")
    return
  }
  
  for _, c := range s.Candidates {
    if c.Number == number {
      util.SetErrMsg("There is already a candidate with this number!")
      return
    }
  }
  
  if len(number) != s.NumberLength {
    util.SetErrMsg("The candidate number '" + number + "' doesn't match the section's candidate number length of " + strconv.Itoa(s.NumberLength) + ".")
    return
  }
  
  if number == "" {
    util.SetErrMsg("The candidate number cannot be empty!")
    return
  }
  
  s.Candidates = append(s.Candidates, util.NewCandidate(name, vice, number))
}

func GetCandidateNumber() string {
  fmt.Print("Enter the candidate number: ")
  util.Scanner.Scan()
  
  return util.Scanner.Text()
}