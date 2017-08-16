package main

import "fmt"

func main() {
	// x := 1
	// p := &x
	// *p = 2
	// fmt.Println(x)

	// ss := []string{"zzwind"}
	// pp := &ss[0]
	// fmt.Println(*pp)
	// fmt.Println(ss)

	var m = map[string]int{}
	m["one"] = 1
	m["two"] = 2

	v, ok := m["two"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在的")
	}

}
