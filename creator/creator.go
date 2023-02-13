package creator

import (
  "fmt"
  
  "ballot/util"
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

type Section struct {
  Name string
  Candidates []Candidate
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
    
    fmt.Println("\nSections:\n")
    
    if len(editing.Sections) > 0 {
      for _, s := range editing.Sections {
       fmt.Printf("%s | %d candidates | %d numbers\n", s.Name, len(s.Candidates), s.NumberLength); 
      }
    } else {
      fmt.Println("There are no sections.")
    }
    
    fmt.Println("\nChoose an option:\n")
    
    fmt.Println("1 - Add new section")
    fmt.Println("2 - Edit an existing section")
    
    fmt.Println("0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        AddSection()
        break
      
      case "2":
        
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

// ----

func AddSection() {
  fmt.Print("Enter the section name: ")
  util.Scanner.Scan()
  name := util.Scanner.Text()
  
  fmt.Print("Enter the number length: ")
  util.Scanner.Scan()
  lenStr := util.Scanner.Text()
  
  len, err := strconv.Atoi(lenStr)
  
  if err != nil {
    return
  }
  
  editing.Sections = append(editing.Sections, NewSection(name, len))
}