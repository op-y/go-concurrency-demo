package main

import (
    "context"
    "fmt"
    "time"
)

func HandleRequest(ctx context.Context) {
    for {
        select {
        case <-ctx.Done(): // 这里会hung住无法结束
            fmt.Println("Handle Request Done.")
            return
        default:
            fmt.Println("Handle Request running... ", ctx.Value("param"))
            time.Sleep(2 * time.Second)
        }
    }
}

func main() {
    ctx := context.WithValue(context.Background(), "param", "1")
    go HandleRequest(ctx)

    // 为了演示这里直接sleep，更合理的方法是用WaitGroup/channel等.
    time.Sleep(10 * time.Second)
}
