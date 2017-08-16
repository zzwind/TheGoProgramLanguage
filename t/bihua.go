package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	// fmt.Println()
	fetchBihua("龚")

	var mySlice = []byte{244, 244, 244, 244, 244, 244, 244, 244}
	data := binary.BigEndian.Uint64(mySlice)
	fmt.Println(data)

}

func fetchBihua(s string) {

	r, err := http.Get("http://www.baidu.com/s?wd=" + s + "%20笔画")
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(r.Body)
	re := regexp.MustCompile(`<div class="op_exactqa_detail_s_answer">(?:[.\s]*)<span>(\d*)</span>(?:[.\s]*)</div>`)
	bb := re.FindSubmatch(b)


	fmt.Printf("%s",bb[1])
	fmt.Println(bb[1])

	// fmt.Println(binary.BigEndian.Uint32(bb[1]))

	fmt.Println("done")
}
