package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Currency struct {
	Usd struct {
		Code        string  `json:"code"`
		Name        string  `json:"name"`
		InverseRate float64 `json:"inverseRate"`
	} `json:"usd"`
	Eur struct {
		Code        string  `json:"code"`
		Name        string  `json:"name"`
		InverseRate float64 `json:"inverseRate"`
	} `json:"eur"`
	Gbp struct {
		Code        string  `json:"code"`
		Name        string  `json:"name"`
		InverseRate float64 `json:"inverseRate"`
	} `json:"gbp"`
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	url := "http://www.floatrates.com/daily/try.json"

	r, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	var currency Currency

	if err := json.Unmarshal(body, &currency); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(PrettyPrint(currency))

}
