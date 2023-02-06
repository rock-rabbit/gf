package main

import (
	"encoding/json"
	"fmt"
	"github.com/rock-rabbit/gf/container/gmap"
)

func main() {
	m := gmap.Map{}
	s := []byte(`{"name":"john","score":100}`)
	json.Unmarshal(s, &m)
	fmt.Println(m.Map())
}
