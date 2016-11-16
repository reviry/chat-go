package main

import (
  "fmt"
  "net"
)

func main() {
  listener, err := net.Listen("tcp", "localhost:8888")
  if err != nil {
    fmt.Printf("Listen error: %s\n", err)
    return
  }
  defer listener.Close()

  conn, err := listener.Accept()
  if err != nil {
    fmt.Printf("Accept error: %s\n", err)
    return
  }
  defer conn.Close()

  fmt.Println("You received a new message!")
  buf := make([]byte, 1024)
  for {
    n, err := conn.Read(buf)
    if n == 0 {
      break
    }
    if err != nil {
      fmt.Printf("Read error: %s\n", err)
    }
    fmt.Print(string(buf[:n]))
  }
}
