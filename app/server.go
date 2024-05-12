package main

import (
    "fmt"
    "github.com/codecrafters-io/redis-starter-go/app/engine"
)

func main() {
    // You can use print statements as follows for debugging, they'll be visible when running tests.
    fmt.Println("Logs from your program will appear here!")

    // Uncomment this block to pass the first stage
    //
    engine.StartEngine(
        engine.RedisConf{
            Host: "0.0.0.0",
            Port: "6379",
        },
    )
}
