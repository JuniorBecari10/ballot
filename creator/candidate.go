package creator

import (
	"fmt"

	"ballot/util"
	"strconv"
)

func EditCandidate(s *util.Section, c *util.Candidate) {
  for {
    util.Clear()
    util.PrintName()
    fmt.Printf("Editing ballot box %s / section %s / candidate %s\n", Editing.Name, s.Name, c.Name)
    util.PrintErrMsg()
    
    fmt.Println("\nChoose a field to edit:\n")
    
    fmt.Printf("1 - Name: %s\n", c.Name)
    fmt.Printf("2 - Vice: %s\n", c.Vice)
    fmt.Printf("3 - Number: %s\n", c.Number)
    
    fmt.Println("\n0 - Go back")
    
    fmt.Print("\n> ")
    util.Scanner.Scan()
    op := util.Scanner.Text()
    
    switch op {
      case "1":
        fmt.Print("Enter the new candidate name: ")
        util.Scanner.Scan()
        
        name := util.Scanner.Text()
        
        if name == "" {
          util.SetErrMsg("The candidate name cannot be empty!")
          break
        }
        
        c.Name = name
        
        break
      
      case "2":
        fmt.Print("Enter the new candidate vice: ")
        util.Scanner.Scan()
        
        vice := util.Scanner.Text()
        c.Vice = vice
        break
      
      case "3":
        fmt.Print("Enter the new candidate number: ")
        util.Scanner.Scan()
        
        number := util.Scanner.Text()
        
        _, err := strconv.Atoi(number)
  
        if err != nil {
          util.SetErrMsg("Couldn't process candidate number: '" + number + "'.")
          break
        }
        
        br := false
        
        for _, c := range s.Candidates {
          if c.Number == number {
            br = true
            util.SetErrMsg("There is already a candidate with this number!")
            break
          }
        }
        
        if br {
          break
        }
        
        if len(number) != s.NumberLength {
          util.SetErrMsg("The candidate number '" + number + "' doesn't match the section's candidate number length of " + strconv.Itoa(s.NumberLength) + ".")
          break
        }
        
        if number == "" {
          util.SetErrMsg("The candidate number cannot be empty!")
          return
        }
        
        c.Number = number
        break
      
      case "0":
        return
    }
  }
}