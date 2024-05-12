package engine

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
)


type RedisConf struct {
    Host string
    Port string
}

func StartEngine(conf RedisConf) {
    l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Host, conf.Port))
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

    listener := make(chan os.Signal, 1)
    signal.Notify(listener, os.Interrupt)
    go func(){
        <-listener
        fmt.Printf("Shutting down redis\n")
        os.Exit(0)
    }()
    for {
        connectionReader(c)
    }

}


func connectionReader(c net.Conn) {
    var buf = make([]byte, 1024)
    _, err := c.Read(buf)
    if err != nil {
        fmt.Println("Failed to read into buffer")
        os.Exit(1)
    }
    for s, val := range strings.Split(string(buf), "\r\n") {
        fmt.Printf("Result : %d %s\n", s, val) 
        if val == "PING" {
            _, err = c.Write([]byte("+PONG\r\n"))
        }

        if err != nil {
            fmt.Println("Failed to write into buffer")
            os.Exit(1)
        }
    }
}
