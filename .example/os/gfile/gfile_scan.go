package main

import (
	"github.com/rock-rabbit/gf/os/gfile"
	"github.com/rock-rabbit/gf/util/gutil"
)

func main() {
	gutil.Dump(gfile.ScanDir("/Users/john/Documents", "*.*"))
	gutil.Dump(gfile.ScanDir("/home/john/temp/newproject", "*", true))
}
