package main

import (
  "fmt"
	"flag"
	"github.com/umanoda/sokudoku"

  "os"
)

var (
  wait = flag.Int("w", 250, "Wait time.") // Intervel time (ms) par display next word.
)

func main() {
  flag.Parse()

  err := sokudoku.Run(*wait)
  if err != nil {
    fmt.Println("Error occered : ", err)
    os.Exit(1)
  }
  os.Exit(0)
}
