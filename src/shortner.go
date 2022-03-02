package main

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

type BitlyResponse struct{
	Link string `json:"link"`
}



func shortenURL(url string) (string,error) {
	client := resty.New();
	resp,err := client.R().SetHeader("Authorization","Bearer "+ACCESSTOKEN).SetHeader("Content-Type","application/json").SetBody(map[string]interface{}{"long_url":url}).Post("https://api-ssl.bitly.com/v4/shorten");
	bitlyResponse := BitlyResponse{};
	json.Unmarshal([]byte(resp.String()),&bitlyResponse)
	return bitlyResponse.Link,err
}