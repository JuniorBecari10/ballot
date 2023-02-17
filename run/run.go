package run

import (
  "fmt"
  
  "ballot/util"
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
    
    fmt.Println("Candidates:")
    
    if b.Config.ShowCandList {
      for _, c := range s.Candidates {
        fmt.Printf("%s | vice: %s | number: %s\n", c.Name, c.Vice, c.Number);
      }
    }
    
    if b.Config.AllowBlank {
      fmt.Println("\nYou're allowed to vote blank in this election. Just press Enter.")
    }
    
    if b.Config.AllowNull {
      fmt.Println("\nYou're allowed to vote null in this election. Just type a number of a candidate that doesn't exist inside this section.")
    }
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    
    nb := util.Scanner.Text()
    var c *util.Candidate = nil
    
    for _, cd := range s.Candidates {
      if nb == cd.Number {
        c = cd
      }
    }
    
    if nb == "" {
      if b.Config.AllowBlank {
        blankVotes++
        sectionIndex++
      } else {
        util.SetErrMsg("You cannot vote blank!")
      }
      
      continue
    }
    
    if c == nil {
      if b.Config.AllowNull {
        nullVotes++
        sectionIndex++
      } else {
        util.SetErrMsg("You cannot vote null!")
      }
      
      continue
    }
     
    // reaching here means that the user has voted in an valid candidate
    c.Votes++
  }
}