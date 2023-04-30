package main

import (
	"luago/state"
)

func main() {
	ls := state.New()
	ls.OpenLibs()
	ls.LoadFile("luac.out")
	ls.Call(0, -1)
}
