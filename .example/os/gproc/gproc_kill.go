package main

import (
	"fmt"
	"github.com/rock-rabbit/gf/os/gproc"
)

func main() {
	pid := 32556
	m := gproc.NewManager()
	m.AddProcess(pid)
	err := m.KillAll()
	fmt.Println(err)
	m.WaitAll()
	fmt.Printf("%d was killed\n", pid)
}
