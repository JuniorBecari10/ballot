package util

import (
  "bufio"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "runtime"
  "strings"
  "strconv"
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

func BoolToYes(b bool) string {
  if b {
    return "Yes"
  } else {
    return "No"
  }
}

func GetOrdinal(n int) string {
  str := strconv.Itoa(n)
  s := string(str[len(str) - 1])
  
  switch s {
    case "1":
      return "st"
    
    case "2":
      return "nd"
    
    case "3":
      return "rd"
    
    default:
      return "th"
  }
}

func Confirm() bool {
  fmt.Print("Are you sure? (y/n) ")
  
  Scanner.Scan()
  
  return strings.ToLower(Scanner.Text()) == "y"
}

func ConfirmBlank() bool {
  fmt.Println("You're going to vote blank.")
  fmt.Print("Do you confirm? (y/n) ")
  
  Scanner.Scan()
  
  return strings.ToLower(Scanner.Text()) == "y"
}

func ConfirmNull() bool {
  fmt.Println("You're going to vote null.")
  fmt.Print("Do you confirm? (y/n) ")
  
  Scanner.Scan()
  
  return strings.ToLower(Scanner.Text()) == "y"
}

func ConfirmCand(c *Candidate) bool {
  fmt.Printf("You're going to vote for the candidate %s.\n", c.Name)
  fmt.Print("Do you confirm? (y/n) ")
  
  Scanner.Scan()
  
  return strings.ToLower(Scanner.Text()) == "y"
}

//func LoadBallot(filename string) *Ballot {
//
//}

// ---

type Ballot struct {
  Name     string     `json:"Name"`
  Sections []*Section `json:"Sections"`

  Config BallotConfig `json:"Config"`
}

func NewBallot(name string) *Ballot {
  return &Ballot { Name: name, Config: BallotConfig {AllowNull: true, AllowBlank: true, ShowCandList: true } }
}

type BallotConfig struct {
  AllowNull    bool
  AllowBlank   bool
  ShowCandList bool
}

// ---

type Section struct {
  Name         string
  Candidates   []*Candidate
  NumberLength int
}

func NewSection(name string, numberLen int) *Section {
  return &Section { Name: name, NumberLength: numberLen }
}

// ---

type Candidate struct {
  Name   string
  Vice   string

  Number string // yes, a string. Because you can set "05" as a number.

  Votes  int // it's necessary.
}

func NewCandidate(name string, vice string, number string) *Candidate {
  return &Candidate { Name: name, Vice: vice, Number: number }
}