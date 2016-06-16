package main

import (
  "fmt"
	"flag"
	"github.com/umanoda/sokudoku"

  "os"
)

func main() {
  wait := flag.Int("w", 250, "Wait time.")
  flag.Parse()

  err := sokudoku.Run(*wait)
  if err != nil {
    fmt.Println("Error occered : ", err)
    os.Exit(1)
  }
  os.Exit(0)
}
