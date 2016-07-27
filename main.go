package main

import (
	_ "github.com/chai2010/cgo-bug-20160727/a"
	_ "github.com/chai2010/cgo-bug-20160727/b"
)

func main() {
	println("cgo")
}
