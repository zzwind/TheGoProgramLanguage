package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separtor")

var pp = flag.Bool("pp", true, "print name")

func main() {
	fmt.Println(*pp)
	flag.Parse()
	fmt.Printf(strings.Join(flag.Args(), *sep))
	fmt.Println(*n)
	if !*n {
		fmt.Println()
		fmt.Println("o ye")
	}
	fmt.Println(*pp)
	if *pp {
		fmt.Println("zzwind")
	}

}
