package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func OnPage(link string) string {
	res, err := http.Get(link)

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func main() {
	fmt.Println(OnPage("https://ikman.lk/en/ad/suzuki-wagon-r-stingray-2015-for-sale-colombo-1416"))
}
