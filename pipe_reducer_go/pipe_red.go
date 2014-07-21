package main

import (
  "os"
  "bufio"
  "fmt"
  "sort"
  "math"
)

type UrlHash struct {
  Url string
  Count int
}
type ByCount []UrlHash

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

func main() {
  path_hash := map[string]int{}

  scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
  for scanner.Scan() {
    path_hash[scanner.Text()] = path_hash[scanner.Text()] + 1
  }

  keys := []UrlHash{}

  for key := range path_hash {
    keys = append(keys, UrlHash{key, path_hash[key]})
  }
  sort.Sort(ByCount(keys))

  min := int(math.Min(float64(len(keys)), 20))

  for k := range keys[0:min] {
    fmt.Println(keys[k].Count, keys[k].Url)
  }
}
