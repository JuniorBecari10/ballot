package creator

import (
  "fmt"
  
  "ballot/util"
  "strings"
  "strconv"
)

var editing *Ballot

type Ballot struct {
  Name string
  Sections []*Section  
}

func NewBallot(name string) *Ballot {
  return &Ballot { Name: name }
}

func CreateMenu() bool {
  fmt.Print("Enter the ballot name: ")
  util.Scanner.Scan()
  
  name := util.Scanner.Text()
  
  if name == "" {
    util.SetErrMsg("The ballot name cannot be empty!")
    return false
  }
  
  editing = NewBallot(name)
  return true
}

func MainMenu() {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot %s\n", editing.Name)
    util.PrintErrMsg()
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Edit sections")
    fmt.Println("2 - Edit ballot name")
    
    fmt.Println("0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        EditSections()
        break
      
      case "2":
        EditBallotName()
        break
      
      case "0":
        return
    }
  }
}

func EditSections() {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot %s / sections\n", editing.Name)
    util.PrintErrMsg()
    
    fmt.Println("\nSections:\n")
    
    if len(editing.Sections) > 0 {
      for _, s := range editing.Sections {
       fmt.Printf("%s | candidates: %d | candidate number length: %d\n", s.Name, len(s.Candidates), s.NumberLength); 
      }
    } else {
      fmt.Println("There are no sections.")
    }
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Add new section")
    
    if len(editing.Sections) > 0 {
      fmt.Println("2 - Edit an existing section")
    }
    
    fmt.Println("0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        AddSection()
        break
      
      case "2":
        if len(editing.Sections) <= 0 {
          break
        }
        
        name := GetSectionName()
        var s *Section = nil
        
        for _, sc := range editing.Sections {
          if strings.ToLower(sc.Name) == strings.ToLower(name) {
            s = sc
          }
        }
        
        if s == nil {
          break
        }
        
        EditSection(s)
        break
      
      case "0":
        return
    }
  }
}

func EditBallotName() {
  fmt.Print("Enter the new ballot name: ")
  util.Scanner.Scan()
  
  name := util.Scanner.Text()
  
  if name == "" {
    util.SetErrMsg("The ballot name cannot be empty!")
    return
  }
  
  editing.Name = name
}

// ----

func AddSection() {
  fmt.Print("Enter the section name: ")
  util.Scanner.Scan()
  name := util.Scanner.Text()
  
  if name == "" {
    util.SetErrMsg("The section name cannot be empty!")
    return
  }
  
  for _, s := range editing.Sections {
    if strings.ToLower(s.Name) == strings.ToLower(name) {
      util.SetErrMsg("There is already a section with this name!")
      return
    }
  }
  
  fmt.Print("Enter the number length: ")
  util.Scanner.Scan()
  lenStr := util.Scanner.Text()
  
  len, err := strconv.Atoi(lenStr)
  
  if err != nil {
    util.SetErrMsg("Couldn't process candidate number length: '" + lenStr + "'.")
    return
  }
  
  editing.Sections = append(editing.Sections, NewSection(name, len))
}

func GetSectionName() string {
  fmt.Print("Enter the section name: ")
  util.Scanner.Scan()
  
  return util.Scanner.Text()
}