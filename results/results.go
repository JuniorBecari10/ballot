package results

import (
  "fmt"
  "sort"
  
  "ballot/util"
)

func ShowResults(b *util.Ballot, blank, null int) {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Viewing results of the election for the ballot box %s\n", b.Name)
    util.PrintErrMsg()
    
    fmt.Println("\nSections (most voted first):\n")
    
    for _, s := range b.Sections {
      sort.Slice(s.Candidates, func (i, j int) bool {
        return s.Candidates[i].Votes > s.Candidates[j].Votes
      })
      
      fmt.Println("section " + s.Name + ":")
      
      for i, c := range s.Candidates {
        fmt.Printf("%d%s - %s | %d votes\n", i + 1, util.GetOrdinal(i + 1), c.Name, c.Votes)
      }
      
      fmt.Println()
      
      if b.Config.AllowBlank {
        fmt.Printf("%d blank vote(s).\n", blank)
      }
      
      if b.Config.AllowNull {
        fmt.Printf("%d null vote(s).\n", null)
      }
      
      fmt.Println()
    }
    
    // TODO: save ballot box and add option to save election results
    fmt.Println("Press Enter to return.")
    util.Scanner.Scan()
    
    return
  }
}