package util

import (
  "fmt"
  "runtime"
  "os/exec"
  "os"
  "bufio"
  "encoding/json"
  "io/ioutil"
)

var errMsg string
var Scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func Clear() {
  switch runtime.GOOS {
    case "linux":
      cmd := exec.Command("clear")
      cmd.Stdout = os.Stdout
      cmd.Run()
      break
    
    case "windows":
      cmd := exec.Command("cmd", "/c", "cls")
      cmd.Stdout = os.Stdout
      cmd.Run()
      break
  }
}

func PrintName() {
  fmt.Println("Ballot Box Creator")
}

func SetErrMsg(msg string) {
  errMsg = msg
}

func PrintErrMsg() {
  if errMsg == "" {
    return
  }
  
  fmt.Print("ERROR: ")
  fmt.Println(errMsg)
  errMsg = ""
}

func SaveBallot(b *Ballot) {
  content, err := json.Marshal(b)
  
  if err != nil {
    panic(err)
  }
  
  err = ioutil.WriteFile(b.Name + ".bb", content, 0777) // ModePerm
  
  if err != nil {
    panic(err)
  }
}

//func LoadBallot(filename string) *Ballot {
//  
//}

// ---

type Ballot struct {
  Name     string     `json:"Name"`
  Sections []*Section `json:Sections`
}

func NewBallot(name string) *Ballot {
  return &Ballot { Name: name }
}

// ---

type Section struct {
  Name string
  Candidates []*Candidate
  NumberLength int
}

func NewSection(name string, numberLen int) *Section {
  return &Section { Name: name, NumberLength: numberLen }
}

// ---

type Candidate struct {
  Name string
  Vice string
  
  Number string // yes, a string. Because you can set "05" as a number.
}

func NewCandidate(name string, vice string, number string) *Candidate {
  return &Candidate { Name: name, Vice: vice, Number: number }
}