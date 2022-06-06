package main

import (
	"fmt"
	"github.com/WhianGaulin/myimooc/fuck"
	"github.com/WhianGaulin/myimooc/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fuckFromWZY()
	fmt.Println(fuck.Fuckwho("you"))
}

func fuckFromWZY() {
	fmt.Println("fuck you hey yo!")
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
