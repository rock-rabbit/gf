package main

import (
	"fmt"

	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/util/gconv"
)

func main() {
	conn := g.Redis().Conn()
	defer conn.Close()
	conn.Do("SET", "k", "v")
	v, _ := conn.Do("GET", "k")
	fmt.Println(gconv.String(v))
}
