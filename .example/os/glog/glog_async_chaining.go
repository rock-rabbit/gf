package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		g.Log().Async().Print("async log", i)
	}
	time.Sleep(time.Second)
}
