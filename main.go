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
    fmt.Println("Ballot Box Creator")
    
    sc.Scan()
  }
}