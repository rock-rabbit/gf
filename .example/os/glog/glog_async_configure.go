package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"time"
)

func main() {
	g.Log().SetAsync(true)
	for i := 0; i < 10; i++ {
		g.Log().Print("async log", i)
	}
	time.Sleep(time.Second)
}
