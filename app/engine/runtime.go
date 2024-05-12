package engine

import (
    "fmt"
    "net"
    "os"
    "io"
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

    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting connection: ", err.Error())
            os.Exit(1)
        }

        go connectionReader(c)
    }
}


func connectionReader(c net.Conn) {
    defer c.Close()

    var buf = make([]byte, 1024)
    for {
        _, err := c.Read(buf)
        if err == io.EOF {
            c.Close()
        }
        if err != nil {
            fmt.Println("Failed to read into buffer")
            return
        }
        _, err = c.Write([]byte("+PONG\r\n"))
        if err != nil {
            fmt.Println("Failed to write into buffer")
            return
        }
    }

}
