package web

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

func WebTwitterStatusMain() {
	resp, err := http.Get("http://twitter.com/users/Goodland.xml")
	checkTwitterError(err)
	user := User{xml.Name{"", "user"}, Status{""}}
	var buf []byte
	buf, err = ioutil.ReadAll(resp.Body)
	checkTwitterError(err)
	xml.Unmarshal(buf, &user)
	fmt.Printf("status: %s\n", user.Status.Text)
}

func checkTwitterError(err error) {
	if err != nil {
		panic("error get twitter:" + err.Error())
	}
}
