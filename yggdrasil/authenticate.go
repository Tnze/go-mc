package yggdrasil

import (
	"fmt"
	"github.com/google/uuid"
)

type Access struct {
	ar authResp
	ct string
}

// agent is a struct of auth
type agent struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

var defaultAgent = agent{
	Name:    "Minecraft",
	Version: 1,
}

// authPayload is a yggdrasil request struct
type authPayload struct {
	Agent       agent  `json:"agent"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	ClientToken string `json:"clientToken"`
	RequestUser bool   `json:"requestUser"`
}

// authResp is the response from Mojang's auth server
type authResp struct {
	Error        string `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	Cause        string `json:"cause"`

	AccessToken       string `json:"accessToken"`
	ClientToken       string `json:"clientToken"` // identical to the one received
	AvailableProfiles []struct {
		ID     uuid.UUID `json:"ID"` // hexadecimal
		Name   string    `json:"name"`
		Legacy bool      `json:"legacy"` // In practice, this field only appears in the response if true. Default to false.
	} `json:"availableProfiles"`                  // only present if the agent field was received

	SelectedProfile struct { // only present if the agent field was received
		ID     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		Legacy bool      `json:"legacy"`
	} `json:"selectedProfile"`
	User struct { // only present if requestUser was true in the request authPayload
		ID         uuid.UUID `json:"id"` // hexadecimal
		Properties []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		}
	} `json:"user"`
}

// Authenticate authenticates a user using their password.
func Authenticate(user, password string) (*Access, error) {
	// Payload
	pl := authPayload{
		Agent:       defaultAgent,
		UserName:    user,
		Password:    password,
		ClientToken: uuid.New().String(),
		RequestUser: true,
	}
	// Resp
	var ar authResp

	// Request
	err := post("/authenticate", pl, &ar)
	if err != nil {
		return nil, err
	}

	if ar.Error != "" {
		err = fmt.Errorf("auth fail: %s: %s, %s}",
			ar.Error, ar.ErrorMessage, ar.Cause)
		return nil, err
	}

	return &Access{ar: ar, ct: pl.ClientToken}, nil
}

func (a *Access) SelectedProfile() (ID uuid.UUID, Name string) {
	return a.ar.SelectedProfile.ID, a.ar.SelectedProfile.Name
}

func (a *Access) AccessToken() string {
	return a.ar.AccessToken
}
