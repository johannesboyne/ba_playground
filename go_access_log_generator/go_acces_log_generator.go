package main

import (
  "os"
  "strconv"
  "math/rand"
  "flag"
  "time"
)

type AccessLogEntry struct {
  IP string
  Timestamp string
  Request string
  HTTPReplyCode int
  Bytes int
}

func (a AccessLogEntry) String() string {
  return a.IP +
  " - - [" +
  a.Timestamp +
  "] \"" +
  a.Request +
  " HTTP/1.0\" " +
  strconv.Itoa(a.HTTPReplyCode) +
  " " +
  strconv.Itoa(a.Bytes) +
  "\n"
}

func createRandomLog () string {
  ips := []string{
    "127.0.0.1",
    "182.168.23.2",
    "122.198.18.2",
    "102.128.29.1",
    "122.168.22.1",
    "130.113.2.2"}
  requests := []string{
    "GET /",
    "GET /hello/",
    "GET /hello/world/",
    "GET /test/",
    "GET /test/animated.gif",
    "GET /test/subdir/",
    "GET /test/subdir00/",
    "GET /test/subdir01/",
    "GET /test/subdir02/",
    "GET /test/subdir03/",
    "GET /test/subdir04/",
    "GET /test/subdir05/",
    "GET /test/subdir06/",
    "GET /test/subdir07/",
    "GET /test/subdir08/",
    "GET /test/subdir09/",
    "GET /test/subdir10/",
    "GET /test/subdir11/",
    "GET /test/subdir12/",
    "GET /test/subdir13/",
    "GET /test/subdir14/",
    "GET /test/dj-sound.mp3",
    "GET /test/mypath/check.html"}
  return AccessLogEntry{
    ips[rand.Intn(len(ips))],
    time.Now().Format("02/Jan/2006:15:04:05 -0700"),
    requests[rand.Intn(len(requests))],
    200,
    1234}.String()
}

var maxMegaBytes int

func main() {
  flag.IntVar(&maxMegaBytes, "max-MB", 1, "max megabytes")
  flag.IntVar(&maxMegaBytes, "m", 1, "max megabytes (shorthand)")
  flag.Parse()

  bytes := 0

  for bytes < (maxMegaBytes*1048576) {
    out := createRandomLog()
    bytes += len([]byte(out))
    os.Stdout.WriteString(out)
  }
}
