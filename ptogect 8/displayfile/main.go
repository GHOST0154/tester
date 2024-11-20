package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("File name missing")
        return
    }
    args := os.Args[1]

    if len(os.Args) > 2 {
        fmt.Println("Too many Arguments")
        return
    }
    contents, err := os.ReadFile(args)

    if err != nil {
        fmt.Println("Error", err)
    } 
        fmt.Print(string(contents))
}