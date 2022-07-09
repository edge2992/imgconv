package main

import (
	"flag"
	"fmt"
	"os"

  imgconv "github.com/edge2992/imgconv"
)

var (
	flagTo   = imgconv.PNG
	flagFrom = imgconv.GIF
)

func init() {
	flag.Var(&flagTo, "to", "after format")
	flag.Var(&flagFrom, "from", "before format")
}

func main() {
	if err := imgconv.ConvertAll(os.Args[1], flagFrom, flagTo); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}
