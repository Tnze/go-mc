package yggdrasil

import "fmt"

// Validate checks if an accessToken is usable for authentication with a Minecraft server.
func (a *Access) Validate() (bool, error) {
	pl := struct {
		AccessToken string `json:"accessToken"`
		ClientToken string `json:"clientToken"`
	}{
		AccessToken: a.ar.AccessToken,
		ClientToken: a.ar.ClientToken,
	}

	resp, err := rowPost("/validate", pl)
	if err != nil {
		return false, fmt.Errorf("request fail: %v", err)
	}

	return resp.StatusCode == 204, resp.Body.Close()
}
