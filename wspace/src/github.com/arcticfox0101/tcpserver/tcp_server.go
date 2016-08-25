package main

import (
    "fmt"
    "net"
    //"os"
    "bufio"
    "strings"
    "github.com/arcticfox0101/stringutil"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3337"
    CONN_TYPE = "tcp"
)

func main() {

  fmt.Println("Launching server...")

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":3337")

  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  for {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // output message received
    fmt.Print("Message Received:", string(message))
    // sample process for string received
    newmessage := stringutil.Reverse(message)
    // send new string back to client
    fmt.Print("Message Reversed:", string(strings.TrimSpace(newmessage)))

    message2 := string(strings.TrimSpace(newmessage))
    conn.Write([]byte(message2 + "\n"))

    temp := string(message)
    temps := strings.Split(temp,"\\n")


   for i := range temps {

     r := stringutil.Reverse(temps[i])
     conn.Write([]byte(r))

     conn.Write([]byte("\n"))

   }

  }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  reqLen, err := conn.Read(buf)
//	err := conn.Read(buf)
fmt.Println(string(buf))
  if err != nil {
    fmt.Println("Error reading:", err.Error())
    fmt.Println(reqLen)
  }

  temp := string(buf)
  temps := strings.Split(temp,"\\n")

  //final := ""
  //conn.Write(temps)
 for i := range temps {
   //conn.Write([]byte(temps[i]))
   //conn.Write([]byte("\n"))
   r := stringutil.Reverse(temps[i])
   conn.Write([]byte(r))
   //final := final + r
   //fmt.Printf(final)
   conn.Write([]byte("\n"))
  // conn.Write([]byte(final + "\n"))
 }
  //reversed := stringutil.Reverse(string(buf))
//conn.Write([]byte(final + "\n"))

  //conn.Write([]byte("\n"))
  //conn.Write(reversed) doesnt work needs to written out as bytes

  //conn.Write([]byte(reversed))

    //conn.Write([]byte("\n"))
  //fmt.Printf(stringutil.Reverse("hello, world\nhellowworld\nhello\n"))
  //fmt.Printf("\n\n")
  // Close the connection when you're done with it.
  //conn.Close()
}
