package p1

import (
    "fmt"

    "github.com/op-y/go-concurrency-demo/happpens-before/init/p2"
    "github.com/op-y/go-concurrency-demo/happpens-before/init/trace"
)


var V1_p1 = trace.Trace("init v1_p1", p2.V1_p2)
var V2_p1 = trace.Trace("init v2_p1", p2.V2_p2)

func init() {
    fmt.Println("init func in p1")
}
