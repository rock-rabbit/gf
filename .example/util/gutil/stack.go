package main

import (
	"github.com/gogf/gf/util/gutil"
)

func Test(s *interface{}) {
	//debug.PrintStack()
	gutil.PrintStack()
}

func main() {
	Test(nil)
}
