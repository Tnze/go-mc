package realms

import "fmt"

// Invite invite player to Realm
func (r *Realms) Invite(s Server, name, uuid string) error {
	pl := struct {
		Name string `json:"name"`
		UUID string `json:"uuid"`
	}{Name: name, UUID: uuid}

	return r.post(fmt.Sprintf("/invites/%d", s.ID), pl, struct{}{})
}
