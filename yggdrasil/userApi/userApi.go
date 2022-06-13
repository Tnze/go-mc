package userApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ServicesURL = "https://api.minecraftservices.com"

var client = http.DefaultClient

func post(endpoint string, accessToken string, resp interface{}) error {
	rowResp, err := rawPost(endpoint, accessToken)
	if err != nil {
		return fmt.Errorf("request fail: %v", err)
	}
	defer rowResp.Body.Close()
	err = json.NewDecoder(rowResp.Body).Decode(resp)
	if err != nil {
		return fmt.Errorf("parse resp fail: %v", err)
	}

	return nil
}

func rawPost(endpoint string, accessToken string) (*http.Response, error) {
	PostRequest, err := http.NewRequest(
		http.MethodPost,
		ServicesURL+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("make request error: %v", err)
	}

	PostRequest.Header.Set("Authorization", "Bearer "+accessToken)
	PostRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	// Do
	return client.Do(PostRequest)
}
