package creator

import (
	"fmt"
	"strconv"
	"strings"

	"ballot/run"
	"ballot/util"
)

var Editing *util.Ballot

func CreateMenu() bool {
  fmt.Print("Enter the ballot box name: ")
  util.Scanner.Scan()
  
  name := util.Scanner.Text()
  
  if name == "" {
    util.SetErrMsg("The ballot name cannot be empty!")
    return false
  }
  
  Editing = util.NewBallot(name)
  return true
}

func MainMenu() {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot box %s\n", Editing.Name)
    util.PrintErrMsg()
    //util.SaveBallot(Editing)
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Edit sections")
    fmt.Println("2 - Edit ballot box name")
    
    if len(Editing.Sections) > 0 {
      fmt.Println("3 - Run election")
    }
    
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
        
      case "3":
        if len(Editing.Sections) > 0 {
          run.RunElection(Editing)
        }
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
    fmt.Printf("Editing ballot box %s / sections\n", Editing.Name)
    util.PrintErrMsg()
    
    fmt.Println("\nSections:\n")
    
    if len(Editing.Sections) > 0 {
      for _, s := range Editing.Sections {
       fmt.Printf("%s | candidates: %d | candidate number length: %d\n", s.Name, len(s.Candidates), s.NumberLength); 
      }
    } else {
      fmt.Println("There are no sections.")
    }
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Add new section")
    
    if len(Editing.Sections) > 0 {
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
        if len(Editing.Sections) <= 0 {
          break
        }
        
        name := GetSectionName()
        var s *util.Section = nil
        
        for _, sc := range Editing.Sections {
          if strings.ToLower(sc.Name) == strings.ToLower(name) {
            s = sc
          }
        }
        
        if s == nil {
          util.SetErrMsg("Couldn't find a section with the name '" + name + "'.")
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
  fmt.Print("Enter the new ballot box name: ")
  util.Scanner.Scan()
  
  name := util.Scanner.Text()
  
  if name == "" {
    util.SetErrMsg("The ballot box name cannot be empty!")
    return
  }
  
  Editing.Name = name
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
  
  for _, s := range Editing.Sections {
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
  
  Editing.Sections = append(Editing.Sections, util.NewSection(name, len))
}

func GetSectionName() string {
  fmt.Print("Enter the section name: ")
  util.Scanner.Scan()
  
  return util.Scanner.Text()
}