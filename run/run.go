package run

import (
  "fmt"
  
  "ballot/util"
)

func RunElection(b *util.Ballot) {
  sectionIndex := 0
  
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Running election for ballot box %s\n", b.Name)
    util.PrintErrMsg()
    
    for {
      s := b.Sections[sectionIndex]
      
      fmt.Printf("\nVoting for section %s:\n\n", s.Name)
      
      fmt.Println("Candidates:")
      
      if b.Config.ShowCandList {
        for _, c := range s.Candidates {
          fmt.Printf("%s | vice: %s | number: %s\n", c.Name, c.Vice, c.Number);
        }
      }
      
      if b.Config.AllowBlank {
        fmt.Println("\nYou're allowed to vote blank in this election. Just press Enter.\n")
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
      
      
    }
  }
}