package main

import (
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    // You can use print statements as follows for debugging, they'll be visible when running tests.
    fmt.Println("Logs from your program will appear here!")

    // Uncomment this block to pass the first stage
    //
    l, err := net.Listen("tcp", "0.0.0.0:6379")
    if err != nil {
        fmt.Println("Failed to bind to port 6379")
        os.Exit(1)
    }
    c, err := l.Accept()
    if err != nil {
        fmt.Println("Error accepting connection: ", err.Error())
        os.Exit(1)
    }

    defer c.Close()

    var buf = make([]byte, 128)

    _, err = c.Read(buf)
    if err != nil {
        fmt.Println("Failed to read into buffer")
        os.Exit(1)
    }
    for s := range strings.Split(string(buf), "\r\n") {
        fmt.Printf("Result : %d\n", s) 

        _, err = c.Write([]byte("+PONG\r\n"))

        if err != nil {
            fmt.Println("Failed to write into buffer")
            os.Exit(1)
        }
    }
}
