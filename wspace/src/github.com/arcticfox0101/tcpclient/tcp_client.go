package main

import "net"
import "fmt"
import "bufio"
import "os"
//import "io"
//import "log"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "localhost:3337")
  for {
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply

    message, _ := bufio.NewReader(conn).ReadString('\n')

     //reply := make([]byte, 1024)
    // Read the incoming connection into the buffer.




      // println("reply from server=", string(reply))
       // Listen on TCP port 2000 on all interfaces.

    fmt.Print("Message from server: "+message)
  }
}



//echo -n "cat\ndog\nfish" | nc localhost 3334
