package engine

import (
	"fmt"
	"net"
	"os"
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

    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting connection: ", err.Error())
            os.Exit(1)
        }

        defer c.Close()
        go connectionReader(c)
    }
}


func connectionReader(c net.Conn) {

    for {
        var buf = make([]byte, 1024)
        _, err := c.Read(buf)
        if err != nil {
            fmt.Println("Failed to read into buffer")
            os.Exit(1)
        }
        var res = "+"
        for s, val := range strings.Split(string(buf), "\r\n") {
            fmt.Printf("Result : %d %s\n", s, val) 
            if strings.Contains(val, "PING"){
                res += "PONG\r\n"
            }
        }
        _, err = c.Write([]byte(res))
        if err != nil {
            fmt.Println("Failed to write into buffer")
            os.Exit(1)
        }

    }

}
