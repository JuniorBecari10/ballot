package results

import (
  "fmt"
  "sort"
  "os"
  "bufio"
  
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
    
    if util.ConfirmMsg("Do you want to write the results in a file? (y/n) ") {
      fmt.Print("Enter the file name: ")
      
      util.Scanner.Scan()
      PrintResults(util.Scanner.Text(), blank, null, b)
    }
    
    return
  }
}

func PrintResults(filename string, blank, null int, b *util.Ballot) {
  file, err := os.Create(filename)
  
  if err != nil {
    panic(err)
  }
  
  defer file.Close()
  w := bufio.NewWriter(file)
  
  fmt.Println("\nSections (most voted first):\n")
  
  for _, s := range b.Sections {
    sort.Slice(s.Candidates, func (i, j int) bool {
      return s.Candidates[i].Votes > s.Candidates[j].Votes
    })
    
    fmt.Fprintln(w, "section " + s.Name + ":")
    
    for i, c := range s.Candidates {
      fmt.Fprintf(w, "%d%s - %s | %d votes\n", i + 1, util.GetOrdinal(i + 1), c.Name, c.Votes)
    }// reset votes
    
    fmt.Println()
    
    if b.Config.AllowBlank {
      fmt.Fprintf(w, "%d blank vote(s).\n", blank)
    }
    
    if b.Config.AllowNull {
      fmt.Fprintf(w, "%d null vote(s).\n", null)
    }
    
    fmt.Fprintln(w)
  }
  
  w.Flush()
}