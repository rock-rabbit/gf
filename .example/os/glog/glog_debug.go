package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"time"

	"github.com/rock-rabbit/gf/os/gtime"
	"github.com/rock-rabbit/gf/os/gtimer"
)

func main() {
	gtimer.SetTimeout(3*time.Second, func() {
		g.Log().SetDebug(false)
	})
	for {
		g.Log().Debug(gtime.Datetime())
		time.Sleep(time.Second)
	}
}
