package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  sc := bufio.NewScanner(os.Stdin)
  
  for {
    Clear()
    fmt.Println("Ballot Box Creator\n")
    
    fmt.Println("Choose an option:\n")
    
    fmt.Println("1 - Create new ballot box")
    fmt.Println("2 - Open an existing ballot box")
    
    fmt.Println("0 - Exit")
    
    fmt.Print("\n> ")
    sc.Scan()
    op := sc.Text()
    
    switch op {
      case "1":
        break
      
      case "2":
        break
      
      case "0":
        Clear()
        os.Exit(0)
    }
  }
}