package main

import (
  "os"
  "bufio"
  "strings"
)

func main() {
  scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
  for scanner.Scan() {
    slices := strings.Split(scanner.Text(), " ")
    if len(slices) >= 6 {
      os.Stdout.WriteString(slices[6] + "\n")
    }
  }
}
