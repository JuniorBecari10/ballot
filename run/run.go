package run

import (
  "fmt"
  
  "ballot/util"
)

func RunElection(b *util.Ballot) {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Running election for ballot box %s\n", b.Name)
    util.PrintErrMsg()
    
    for _, s := range b.Sections {
      fmt.Printf("\nVoting for section %s:\n\n", s.Name)
      
      fmt.Println("Candidates:")
      
      for _, c := range s.Candidates {
        fmt.Printf("%s | vice: %s | number: %s\n", c.Name, c.Vice, c.Number);
      }
      
      fmt.Print("\n> ")
      util.Scanner.Scan()
      
      nb := util.Scanner.Text()
      var c *Candidate = nil
      
      for _, cd := range s.Candidates {
        if nb == c.Number {
          c = cd
        }
      }
      
      
    }
  }
}