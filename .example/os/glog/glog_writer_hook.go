package main

import (
	"fmt"
	"github.com/rock-rabbit/gf/frame/g"

	"github.com/rock-rabbit/gf/os/glog"
	"github.com/rock-rabbit/gf/text/gregex"
)

type MyWriter struct {
	logger *glog.Logger
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	s := string(p)
	if gregex.IsMatchString(`\[(PANI|FATA)\]`, s) {
		fmt.Println("SERIOUS ISSUE OCCURRED!! I'd better tell monitor in first time!")
		g.Client().PostContent("http://monitor.mydomain.com", s)
	}
	return w.logger.Write(p)
}

func main() {
	glog.SetWriter(&MyWriter{
		logger: glog.New(),
	})
	glog.Debug("DEBUG")
	glog.Fatal("FATAL ERROR")

}
