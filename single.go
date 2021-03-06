package single

import (
   "fmt"
   "net"
   "os"
)

func main() {

  service := ":1200"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  ckerr(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  ckerr(err)

  for {
      conn, err := listener.Accept()
      if err != nil {
          continue
      }

      handleClient(conn)
      conn.Close() // we're finished
}

func handleClient(conn net.Conn) {
    var buf [512]byte
    
    for {

       n, err := conn.Read(buf[0:])
       if err != nil {
            return
       }
       fmt.Println( string(buf[0:]) )
       _, err2 := conn.Write(buf[0:n])
       if err2 != nil {
          return
       }
    }
}

func ckerr(err error) {
    if err != nil {
       fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error() )
       os.Exit(1)
    }
}

