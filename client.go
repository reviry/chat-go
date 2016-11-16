package main

import (
  "fmt"
  "net"
  "bufio"
  "os"
)

func main() {
  conn, err := net.Dial("tcp", "localhost:8888")
  if err != nil {
    fmt.Printf("Dial error: %s\n", err)
    return
  }
  defer conn.Close()

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    sendMsg := scanner.Text() + "\n"
    conn.Write([]byte(sendMsg))
  }
}
