package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func nice_advice() string {
	resp, err := http.Get("http://fucking-great-advice.ru/api/random")
	if err != nil {
		fmt.Println("No response")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	type response struct {
		Id    int    `json:"id"`
		Text  string `json:"text"`
		Sound string `json:"sound"`
	}
	var data response
	json.Unmarshal(body, &data)
	return data.Text

}
