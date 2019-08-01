package authenticate

//Simple yggdrasil-minecraft-login method.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Agent is a struct of auth
type Agent struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

// Payload is a authenticate request struct
type Payload struct {
	Agent       `json:"agent"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	ClientToken string `json:"clientToken"`
	RequestUser bool   `json:"requestUser"`
}

// Allow auth with thirdparty auth server.
func AuthWithThirdPartyServer(user, passwd, AuthServer string) (respData Response, err error) {
	AuthServer += "/authserver/authenticate"
	j, err := json.Marshal(Payload{
		Agent: Agent{
			Name:    "Minecraft",
			Version: 1,
		},
		UserName:    user,
		Password:    passwd,
		ClientToken: "gomcbotauther",
		RequestUser: true,
	})
	// fmt.Println(string(j))
	if err != nil {
		err = fmt.Errorf("encoding json fail: %v", err)
		return
	}

	//Post
	client := http.Client{}
	PostRequest, err := http.NewRequest(http.MethodPost, AuthServer,
		bytes.NewReader(j))
	if err != nil {
		err = fmt.Errorf("make request error: %v", err)
		return
	}
	PostRequest.Header.Set("User-Agent", "go-mc")
	PostRequest.Header.Set("Connection", "keep-alive")
	PostRequest.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(PostRequest)
	if err != nil {
		err = fmt.Errorf("post authenticate fail: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read authenticate resp fail: %v", err)
		return
	}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		err = fmt.Errorf("unmarshal json data fail: %v", err)
		return
	}
	if respData.Error != "" {
		err = fmt.Errorf("authenticate fail: {error: %q, errorMessage: %q, cause: %q}",
			respData.Error, respData.ErrorMessage, respData.Cause)
		return
	}
	return
}

// Authenticate authenticates a user using their password.
func Authenticate(user, passwd string) (respData Response, err error) {
	j, err := json.Marshal(Payload{
		Agent: Agent{
			Name:    "Minecraft",
			Version: 1,
		},
		UserName:    user,
		Password:    passwd,
		ClientToken: "gomcbotauther",
		RequestUser: true,
	})
	// fmt.Println(string(j))
	if err != nil {
		err = fmt.Errorf("encoding json fail: %v", err)
		return
	}

	//Post
	client := http.Client{}
	PostRequest, err := http.NewRequest(http.MethodPost, "https://authserver.mojang.com/authenticate",
		bytes.NewReader(j))
	if err != nil {
		err = fmt.Errorf("make request error: %v", err)
		return
	}
	PostRequest.Header.Set("User-Agent", "go-mc")
	PostRequest.Header.Set("Connection", "keep-alive")
	PostRequest.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(PostRequest)
	if err != nil {
		err = fmt.Errorf("post authenticate fail: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read authenticate resp fail: %v", err)
		return
	}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		err = fmt.Errorf("unmarshal json data fail: %v", err)
		return
	}
	if respData.Error != "" {
		err = fmt.Errorf("authenticate fail: {error: %q, errorMessage: %q, cause: %q}",
			respData.Error, respData.ErrorMessage, respData.Cause)
		return
	}
	return
}

// Response is the response from Mojang's auth server
type Response struct {
	Error        string `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	Cause        string `json:"cause"`

	AccessToken       string `json:"accessToken"`
	ClientToken       string `json:"clientToken"` // identical to the one received
	AvailableProfiles []struct {
		ID     string `json:"ID"` // hexadecimal
		Name   string `json:"name"`
		Legacy bool   `json:"legacy"` // In practice, this field only appears in the response if true. Default to false.
	} `json:"availableProfiles"` // only present if the agent field was received

	SelectedProfile struct { // only present if the agent field was received
		ID     string `json:"id"`
		Name   string `json:"name"`
		Legacy bool   `json:"legacy"`
	} `json:"selectedProfile"`
	User struct { // only present if requestUser was true in the request payload
		ID         string `json:"id"` // hexadecimal
		Properties []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		}
	} `json:"user"`
}
