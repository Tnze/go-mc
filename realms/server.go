package realms

import (
	"errors"
	"fmt"
)

type Server struct {
	ID                   int
	RemoteSubscriptionID string
	Owner                string
	OwnerUUID            string
	Name                 string
	MOTD                 string
	State                string
	DaysLeft             int
	Expired              bool
	ExpiredTrial         bool
	WorldType            string
	Players              []string
	MaxPlayers           int
	MiniGameName         *string
	MiniGameID           *int
	MinigameImage        *string
	ActiveSlot           int
	//Slots                interface{}
	Member bool
}

// Worlds return a list of servers that the user is invited to or owns.
func (r *Realms) Worlds() ([]Server, error) {
	var resp struct {
		Servers []Server
		*Error
	}

	err := r.get("/worlds", &resp)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		err = resp.Error
	}

	return resp.Servers, err
}

// Server returns a single server listing about a server.
// you must be the owner of the server.
func (r *Realms) Server(ID int) (s Server, err error) {
	var resp = struct {
		*Server
		*Error
	}{Server: &s}

	err = r.get(fmt.Sprintf("/worlds/%d", ID), &resp)
	if err != nil {
		return
	}

	if resp.Error != nil {
		err = resp.Error
	}

	return
}

// Address used to get the IP address for a server.
// Call TOS before you call this function.
func (r *Realms) Address(s Server) (string, error) {
	var resp struct {
		Address       string
		PendingUpdate bool

		ResourcePackUrl  *string
		ResourcePackHash *string

		*Error
	}

	err := r.get(fmt.Sprintf("/worlds/v1/%d/join/pc", s.ID), &resp)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		err = resp.Error
		return "", err
	}

	if resp.PendingUpdate {
		return "", errors.New("pending update")
	}
	return resp.Address, err
}

// Backups returns a list of backups for the world.
func (r *Realms) Backups(s Server) ([]int, error) {
	var bs []int
	err := r.get(fmt.Sprintf("/worlds/%d/backups", s.ID), &bs)

	return bs, err
}

//func (r *Realms) Download() (link, resURL, resHash string) {
//	var resp struct {
//		DownloadLink     string
//		ResourcePackURL  *string
//		ResourcePackHash *string
//		*Error
//	}
// TODO: What's the $WORLD(1-4) means?
//	err := r.get(fmt.Sprintf("/worlds/$ID/slot/$WORLD(1-4)/download", s.ID), &resp)
//
//}

// Ops returns a list of operators for this server.
// You must own this server to view this.
func (r *Realms) Ops(s Server) (ops []string, err error) {
	err = r.get(fmt.Sprintf("/ops/%d", s.ID), &ops)
	return
}

// SubscriptionLife returns the current life of a server subscription.
func (r *Realms) SubscriptionLife(s Server) (startDate int64, daysLeft int, Type string, err error) {
	var resp = struct {
		StartDate        *int64
		DaysLeft         *int
		SubscriptionType *string
	}{
		StartDate:        &startDate,
		DaysLeft:         &daysLeft,
		SubscriptionType: &Type,
	}

	err = r.get(fmt.Sprintf("/subscriptions/%d", s.ID), &resp)
	return
}
