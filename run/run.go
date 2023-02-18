package run

import (
  "fmt"
  
  "ballot/util"
  "ballot/results"
)

func RunElection(b *util.Ballot) {
  sectionIndex := 0
  blankVotes := 0
  nullVotes := 0
  
  for {
    if sectionIndex >= len(b.Sections) {
      sectionIndex = 0
    }
    
    util.Clear()
    util.PrintName()
    fmt.Printf("Running election for ballot box %s\n", b.Name)
    util.PrintErrMsg()
    
    s := b.Sections[sectionIndex]
    
    fmt.Printf("\nVoting for section %s:\n\n", s.Name)
    
    if b.Config.ShowCandList {
      fmt.Println("Candidates:\n")
      
      for _, c := range s.Candidates {
        fmt.Printf("%s | vice: %s | number: %s\n", c.Name, c.Vice, c.Number);
      }
    }
    
    fmt.Println()
    
    if b.Config.AllowBlank {
      fmt.Println("You're allowed to vote blank in this election.")
    }
    
    if b.Config.AllowNull {
      fmt.Println("You're allowed to vote null in this election.")
    }
    
    fmt.Println()
    
    fmt.Print("> ")
    util.Scanner.Scan()
    
    nb := util.Scanner.Text()
    var c *util.Candidate = nil
    
    for _, cd := range s.Candidates {
      if nb == cd.Number {
        c = cd
      }
    }
    
    if nb == "exit" {
      if util.Confirm() {
        results.ShowResults(b, blankVotes, nullVotes)
        util.ClearVotes(b)
        return
      }
    }
    
    if nb == "" {
      if b.Config.AllowBlank {
        if util.ConfirmBlank() {
          blankVotes++
          sectionIndex++
        }
      } else {
        util.SetErrMsg("You cannot vote blank!")
      }
      
      continue
    }
    
    if c == nil {
      if b.Config.AllowNull {
        if util.ConfirmNull() {
          nullVotes++
          sectionIndex++
        }
      } else {
        util.SetErrMsg("You cannot vote null!")
      }
      
      continue
    }
     
    // reaching here means that the user has voted in an valid candidate
    if util.ConfirmCand(c) {
      c.Votes++
      sectionIndex++
    }
  }
}