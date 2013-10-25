package main

import (
        "fmt"
        "j4k.co/fmatter"
)

type Metadata struct {
    Title string
    Author string
    Date string
    Tags []string
}

func GetMetadata(data []byte) (*Metadata, []byte) {
    m := Metadata{}
    cont, err := fmatter.Read(data, &m)
    if err != nil {
        fmt.Println("Error parsing metadata:", err)
    }
    return &m, cont
}
