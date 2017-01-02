package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8888")
  checkError(err)
  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)

  defer listener.Close()

  for {
    conn, err := listener.Accept()
    if err != nil {
      continue
    }
    go handleClient(conn)
  }
}

func handleClient(conn net.Conn) {
  defer conn.Close()

  buf := make([]byte, 1024)
  for {
    n, err := conn.Read(buf)
    if n == 0 {
      break
    }
    if err != nil {
      fmt.Printf("Read error: %s\n", err)
    }
    fmt.Println("You received a new message!")
    fmt.Print(string(buf[:n]))
  }
}

func checkError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
