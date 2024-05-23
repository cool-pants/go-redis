package engine

import (
    "fmt"
)

const (
    STRING = "+"
    ERROR = "-"
    INTEGER = ":"
    BULK_STRING = "$"
    ARRAY = "*"
)

type ProtoData struct {
    Type string
    Len uint16
    Data []byte
    DataString string
}

// type RedisType interface {
//     GetValue[]
// }

func ParseData

