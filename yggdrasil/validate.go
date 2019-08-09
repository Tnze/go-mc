package yggdrasil

import (
	"fmt"
	"io/ioutil"
)

// Validate checks if an accessToken is usable for authentication with a Minecraft server.
func (a *Access) Validate() (bool, error) {
	pl := a.ar.tokens

	resp, err := rowPost("/validate", pl)
	if err != nil {
		return false, fmt.Errorf("request fail: %v", err)
	}

	return resp.StatusCode == 204, resp.Body.Close()
}

// Invalidate invalidates accessTokens using a client/access token pair.
func (a *Access) Invalidate() error {
	pl := a.ar.tokens

	resp, err := rowPost("/invalidate", pl)
	if err != nil {
		return fmt.Errorf("request fail: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		content, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("invalidate error: %v: %s", resp.Status, content)
	}

	return nil
}
