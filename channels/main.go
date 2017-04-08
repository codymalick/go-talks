package main

import (
    "time"
    "fmt"
)

func main() {
    // Simple integer channel
    channel := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            channel <- i
            time.Sleep(time.Second)
        }
    }()

    go func() {
        for {
            value := <- channel
            fmt.Println(value)
        }
    }()

    time.Sleep(time.Second * 15)
}