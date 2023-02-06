package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/os/glog"
)

func main() {
	g.Log().SetLevelPrefix(glog.LEVEL_DEBU, "debug")
	g.Log().Debug("test")
}
