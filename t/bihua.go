package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	// fmt.Println()
	for {
		fmt.Println("美女，输入你要查询笔画的词语")

		var name string

		fmt.Scan(&name)

		// name := "龚启舟姜旭园"
		count := 0
		for _, n := range name {
			count += fetchBihua(string(n))
		}
		fmt.Println(name + "笔画总数是:" + strconv.Itoa(count))
		fmt.Println()

		// var mySlice = []byte{244, 244, 244, 244, 244, 244, 244, 244}
		// data := binary.BigEndian.Uint64(mySlice)
		// fmt.Println(data)
	}

}

func fetchBihua(s string) int {

	r, err := http.Get("http://www.baidu.com/s?wd=" + s + "%20笔画")
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(r.Body)
	re := regexp.MustCompile(`<div class="op_exactqa_detail_s_answer">(?:[.\s]*)<span>(\d*)</span>(?:[.\s]*)</div>`)
	bb := re.FindSubmatch(b)

	i, _ := strconv.Atoi(string(bb[1]))
	return i

	// fmt.Println(strconv.Atoi(string(bb[1])))
	// ss, _ := utf8.DecodeRune(bb[1])
	// fmt.Println(utf8.RuneLen(ss))

	// fmt.Printf("%s", bb[1])
	// fmt.Println(bb[1])

	// fmt.Println(binary.BigEndian.Uint16(bb[1]))
	// fmt.Println(binary.LittleEndian.Uint16(bb[1]))

	// fmt.Println("done")
}
