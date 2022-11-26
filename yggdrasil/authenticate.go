package yggdrasil

import (
	"github.com/google/uuid"
)

type Access struct {
	ar authResp
}

// agent is a struct of auth
type agent struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type proof struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Tokens store AccessToken and ClientToken
type Tokens struct {
	AccessToken string `json:"accessToken"`
	ClientToken string `json:"clientToken"`
}

var defaultAgent = agent{
	Name:    "Minecraft",
	Version: 1,
}

// authPayload is a yggdrasil request struct
type authPayload struct {
	Agent agent `json:"agent"`
	proof
	ClientToken string `json:"clientToken,omitempty"`
	RequestUser bool   `json:"requestUser"`
}

// authResp is the response from Mojang's auth server
type authResp struct {
	Tokens
	AvailableProfiles []Profile `json:"availableProfiles"` // only present if the agent field was received

	SelectedProfile Profile `json:"selectedProfile"` // only present if the agent field was received
	User            struct {
		// only present if requestUser was true in the request authPayload
		ID         string `json:"id"` // hexadecimal
		Properties []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		}
	} `json:"user"`

	*Error
}

type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Legacy bool   `json:"legacy"` // we don't care
}

// Authenticate authenticates a user using their password.
func Authenticate(user, password string) (*Access, error) {
	// Payload
	pl := authPayload{
		Agent: defaultAgent,
		proof: proof{
			UserName: user,
			Password: password,
		},
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

	if ar.Error != nil {
		return nil, *ar.Error
	}

	return &Access{ar}, nil
}

func (a *Access) SelectedProfile() (ID, Name string) {
	return a.ar.SelectedProfile.ID, a.ar.SelectedProfile.Name
}

func (a *Access) AccessToken() string {
	return a.ar.AccessToken
}

func (a *Access) AvailableProfiles() []Profile {
	return a.ar.AvailableProfiles
}

func (a *Access) SetTokens(tokens Tokens) {
	a.ar.Tokens = tokens
}

func (a *Access) GetTokens() Tokens {
	return a.ar.Tokens
}
