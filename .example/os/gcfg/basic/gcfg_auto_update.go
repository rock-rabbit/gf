package main

import (
	"fmt"
	"time"

	"github.com/rock-rabbit/gf/os/gtimer"

	"github.com/rock-rabbit/gf/frame/g"
)

// 配置文件热更新示例
func main() {
	c := g.Config()
	// 每隔1秒打印当前配置项值，用户可手动在外部修改文件内容，gcfg读取到的配置项值会即时得到更新
	gtimer.SetInterval(time.Second, func() {
		fmt.Println(c.Get("viewpath"))
	})

	select {}
}
