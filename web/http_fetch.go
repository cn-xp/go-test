package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func WebHttpFetchMain() {
	res, err := http.Get("http://www.baidu.com")
	checkFetchError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkFetchError(err)
	fmt.Printf("Got: %q\n", string(data))
}

func checkFetchError(err error) {
	if err != nil {
		panic("error get: " + err.Error())
	}
}
