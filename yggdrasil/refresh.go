package yggdrasil

import "fmt"

type refreshPayload struct {
	Tokens
	SelectedProfile *Profile `json:"selectedProfile,omitempty"`

	RequestUser bool `json:"requestUser"`
}

// Refresh refreshes a valid accessToken.
//
// It can be used to keep a user logged in between
// gaming sessions and is preferred over storing
// the user's password in a file
func (a *Access) Refresh(profile *Profile) error {
	pl := refreshPayload{
		Tokens:          a.ar.Tokens,
		SelectedProfile: profile, // used to change profile, don't use now
		RequestUser:     true,
	}

	resp := struct {
		*authResp
		*Error
	}{authResp: &a.ar}

	err := post("/refresh", pl, &resp)
	if err != nil {
		return fmt.Errorf("post fail: %v", err)
	}

	if resp.Error != nil {
		return resp.Error
	}

	return nil
}
