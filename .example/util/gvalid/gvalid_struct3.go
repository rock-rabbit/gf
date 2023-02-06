package main

import (
	"context"
	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/util/gvalid"
)

// same校验
func main() {
	type User struct {
		Pass string `gvalid:"passwd1 @required|length:2,20|password3||密码强度不足"`
	}

	user := &User{
		Pass: "1",
	}

	g.Dump(gvalid.CheckStruct(context.TODO(), user, nil).Maps())
}
