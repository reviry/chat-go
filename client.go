package main

import (
  "fmt"
  "net"
  "bufio"
  "os"
)

func main() {
  tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8888")
  checkError(err)
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkError(err)

  defer conn.Close()

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    sendMsg := scanner.Text() + "\n"
    conn.Write([]byte(sendMsg))
  }
}

func checkError(err error) {
  if err != nil {
    fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
    os.Exit(1)
  }
}
