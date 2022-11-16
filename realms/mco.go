package realms

import "io"

// Available returns whether the user can access the Minecraft Realms service
func (r *Realms) Available() (ok bool, err error) {
	err = r.get("/mco/available", &ok)
	return
}

// Compatible returns whether the clients version is up-to-date with Realms.
//
//	if the client is outdated, it returns OUTDATED,
//	if the client is running a snapshot, it returns OTHER,
//	else it returns COMPATIBLE.
func (r *Realms) Compatible() (string, error) {
	resp, err := r.c.Get(Domain + "/mco/client/compatible")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	rp, err := io.ReadAll(resp.Body)

	return string(rp), err
}

// TOS is what to join Realms servers you must agree to.
// Call this function will set this flag.
func (r *Realms) TOS() error {
	resp, err := r.c.Post(Domain+"/mco/tos/agreed", "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
