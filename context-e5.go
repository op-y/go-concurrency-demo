package main

import (
    "context"
    "fmt"
    "time"
)

func HandleRequest(ctx context.Context) {
    go WriteRedis(ctx)
    go WriteDatabase(ctx)

    for {
        select {
        case <-ctx.Done():
            fmt.Println("Handle Request Done.")
            return
        default:
            fmt.Println("Handle Request running...")
            time.Sleep(2 * time.Second)
        }
    }
}

func WriteRedis(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Write Redis Done.")
            return
        default:
            fmt.Println("Write Redis running...")
            time.Sleep(2 * time.Second)
        }
    }
}

func WriteDatabase(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Write Database Done.")
            return
        default:
            fmt.Println("Write Database running...")
            time.Sleep(2 * time.Second)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go HandleRequest(ctx)

    time.Sleep(5 * time.Second)
    fmt.Println("It's time to stop all sub goroutines!")
    cancel()

    // 为了演示这里直接sleep，更合理的方法是用WaitGroup/channel等.
    time.Sleep(5 * time.Second)
}
