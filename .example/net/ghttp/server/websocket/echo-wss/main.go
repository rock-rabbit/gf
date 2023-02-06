package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/net/ghttp"
	"github.com/rock-rabbit/gf/os/gfile"
	"github.com/rock-rabbit/gf/os/glog"
)

func main() {
	s := g.Server()
	s.BindHandler("/wss", func(r *ghttp.Request) {
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(err)
			r.Exit()
		}
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				return
			}
			if err = ws.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})
	s.SetServerRoot(gfile.MainPkgPath())
	s.EnableHTTPS("../../https/server.crt", "../../https/server.key")
	s.SetPort(8199)
	s.Run()
}
