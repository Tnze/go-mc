package yggdrasil

import (
	"encoding/json"
	"fmt"
)

// SignOut invalidates accessTokens using an account's username and password.
func SignOut(user, password string) error {
	pl := proof{
		UserName: user,
		Password: password,
	}

	resp, err := rawPost("/signout", pl)
	if err != nil {
		return fmt.Errorf("request fail: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		var err Error
		if err := json.NewDecoder(resp.Body).Decode(&err); err != nil {
			return fmt.Errorf("unmarshal error fail: %v", err)
		}
		return fmt.Errorf("invalidate error: %v", err)
	}

	return nil
}
