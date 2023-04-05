// Package yggdrasil implement Yggdrasil protocol.
//
// Minecraft 1.6 introduced a new authentication scheme called Yggdrasil
// which completely replaces the previous authentication system.
// Mojang's other game, Scrolls, uses this method of authentication as well.
// Mojang has said that this authentication system should be used by everyone for custom logins,
// but credentials should never be collected from users. ----- https://wiki.vg
package yggdrasil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Err    string `json:"error"`
	ErrMsg string `json:"errorMessage"`
	Cause  string `json:"cause"`
}

func (e Error) Error() string {
	return e.Err + ": " + e.ErrMsg + ", " + e.Cause
}

var AuthURL = "https://authserver.mojang.com"

var client = http.DefaultClient

func post(endpoint string, payload any, resp any) error {
	rowResp, err := rawPost(endpoint, payload)
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

func rawPost(endpoint string, payload any) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload fail: %v", err)
	}

	PostRequest, err := http.NewRequest(
		http.MethodPost,
		AuthURL+endpoint,
		bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("make request error: %v", err)
	}

	PostRequest.Header.Set("User-agent", "go-mc")
	PostRequest.Header.Set("Connection", "keep-alive")
	PostRequest.Header.Set("Content-Type", "application/json")

	// Do
	return client.Do(PostRequest)
}
