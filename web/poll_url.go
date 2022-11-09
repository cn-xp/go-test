package web

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.baidu.com/",
	"http://www.google.com/",
	"http://blog.golang.org/",
}

func WebPollUrlMain() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("error head:", url, err)
			continue
		}

		fmt.Println(url, ":", resp.Status)
	}
}
