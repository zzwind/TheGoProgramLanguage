package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var url string = "https://xueqiu.com/stock/quote_order.json?page=1&size=3000&order=desc&exchange=CN&stockType=sha&column=symbol%2Cname%2Ccurrent%2Cchg%2Cpercent%2Clast_close%2Copen%2Chigh%2Clow%2Cvolume%2Camount%2Cmarket_capital%2Cpe_ttm%2Chigh52w%2Clow52w%2Chasexist&orderBy=percent&_=1491379357605"

//

type Stock struct {
	//Column []string
	Data [][]interface{}
}

//func (s Stock) GET(i, ii int) interface{} {
//
//	return s.Data[i][ii]
//}
var stock = &Stock{}

func main() {

	//for _, v := range stock.Data {
	//	//fmt.Println(v[2])
	//	//fmt.Println(reflect.TypeOf(v))
	//}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	cookie0 := &http.Cookie{Name: "xq_a_token", Value: "653eba279cad4aea82469384444a3d9a2443fd5f"}
	//cookie1 := &http.Cookie{Name: "xq_r_token", Value: "d1accc7b0cafd743be1b975a863a146e514d9c80"}

	req.AddCookie(cookie0)
	//req.AddCookie(cookie1)

	var client = &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Printf("%s", ret)
	err = json.Unmarshal([]byte(ret), stock)

	fmt.Println(err)

	//compare(0.00, 10.00)
	//compare(1.00, 10.00)
	//compare(2.00, 10.00)
	compare(0.1, 10.00)

	//"SH600649","城投控股",14.83,-1.65,-10.01,16.48,0.0,0.0,0.0,0.0,0.0,3.751360665222E10,17.9085,21.41,13.06,"false"
	//fmt.Println(string(ret))

}

func compare(base, now float64) {
	fmt.Println()

	fname := strconv.FormatFloat(base, 'f', -1, 64) + " " + strconv.FormatFloat(now, 'f', -1, 64) + " " + time.Now().Format("2006-01-02 15-04-05")

	f, err := os.Create("stockTxt/" + fname)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString("https://xueqiu.com/S/")
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range stock.Data {

		v1, _ := v[2].(float64)
		v2, _ := v[14].(float64)
		if (v1-v2) <= base && v1 < now {
			//fmt.Println(v[0], v[1], v[2], v[14])

			v0, ok := v[0].(string)
			if !ok {
				fmt.Println(ok)
			}
			v1, ok := v[1].(string)
			if !ok {
				fmt.Println(ok)
			}
			v2, ok := v[2].(string)
			if !ok {
				//fmt.Println(ok)
			}
			v14, ok := v[14].(string)
			if !ok {
				//fmt.Println(ok)
			}
			//_, err = f.WriteString(v1 + " " + "https://xueqiu.com/S/" + v0 + " " + v2 + " " + v14)
			f.WriteString(v1 + " " + "https://xueqiu.com/S/" + v0 + " " + v2 + " " + v14 + "\n")

		}
	}
	f.Close()
}
