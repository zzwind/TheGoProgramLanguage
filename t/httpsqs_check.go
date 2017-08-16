package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

var re = regexp.MustCompile(`"cpc5338"`)

func main() {

	for i := 64550; i < 64614; i++ {
		ResultCheck(i)
		time.Sleep(time.Second * 5)
	}

}

func ResultCheck(s int) {

	curent := strconv.Itoa(s)

	fmt.Println("当前位置：" + curent)

	r, err := http.Get("http://192.168.2.240:1216/?name=ZS_RawMaterialPrice&opt=view&pos=" + curent)
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(r.Body)

	bb := re.Match(b)
	if bb == true {
		fmt.Println("数据找到了，位置：" + curent)
		os.Exit(0)
	}
}
