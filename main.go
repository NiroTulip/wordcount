package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
flag.Parse()
src := strings.Join(flag.Args(), "")
words := strings.Fields(src)
fmt.Println(len(words))
}
