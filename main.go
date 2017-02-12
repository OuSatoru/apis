package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

const BAIDU_SEARCH string = "http://www.baidu.com/s?wd="

func main() {
	http.HandleFunc("/search", searchApi)
	http.ListenAndServe(":7890", nil)
}

func searchApi(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	word := r.Form["wd"][0]
	html := searchBaidu(word)
	fmt.Println(word)
	//fmt.Println(html)
	fmt.Fprint(w, html)
}

func searchBaidu(s string) string {
	html, err := http.Get(BAIDU_SEARCH + s)
	if err != nil {
		panic(err)
	}
	defer html.Body.Close()
	body, _ := ioutil.ReadAll(html.Body)
	return string(body)
}

func teststr(s string) string {
	return "test " + s
}