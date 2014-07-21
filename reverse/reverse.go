package main
import (
  "strings"
  "fmt"
  "bufio"
  "code.google.com/p/rog-go/reverse"
)

func main() {
  const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
  scanner := reverse.NewScanner(strings.NewReader(input))
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}
