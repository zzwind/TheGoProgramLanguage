package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/a", handler3)
	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}

		//cycles, err := strconv.ParseFloat(r.Form["cycles"][0], 64)

		cyclesS := r.FormValue("cycles")
		cycles := 0.00
		if cyclesS != "" {
			var err error
			cycles, err = strconv.ParseFloat(cyclesS, 64)
			if err != nil {
				log.Println(err)
			}
		}

		fmt.Println(cycles)
		lissajous(w, cycles)
	})
	//http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%s]\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	fmt.Fprintf(w, "\n GET \n")
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "\n POST \n")
	for k, v := range r.PostForm {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	fmt.Println(r.Form)

}
