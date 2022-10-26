package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NiceAdvice() string {
	resp, err := http.Get("http://fucking-great-advice.ru/api/random")
	if err != nil {
		fmt.Println("No response")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	type response struct {
		Id    int    `json:"id"`
		Text  string `json:"text"`
		Sound string `json:"sound"`
	}
	var data response
	json.Unmarshal(body, &data)
	return data.Text
}
