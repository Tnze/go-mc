package realms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Realms struct {
	c http.Client
}


type Error struct {
	ErrorCode int
	ErrorMsg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.ErrorCode, e.ErrorMsg)
}

// Domain is the URL of Realms API server
// Panic if it cannot be parse by url.Parse().
var Domain = "https://pc.realms.minecraft.net"

// New create a new Realms c with version, username, accessToken and UUID without dashes.
func New(version, user, astk, uuid string) *Realms {
	r := &Realms{
		c: http.Client{},
	}

	var err error
	r.c.Jar, err = cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	d, err := url.Parse(Domain)
	if err != nil {
		panic("cannot parse realms.Domain as url: " + err.Error())
	}

	r.c.Jar.SetCookies(d, []*http.Cookie{
		{Name: "user", Value: user},
		{Name: "version", Value: version},
		{Name: "sid", Value: "token:" + astk + ":" + uuid},
	})

	return r
}

func (r *Realms) get(endpoint string, resp interface{}) error {
	rawResp, err := r.c.Get(Domain + endpoint)
	if err != nil {
		return err
	}
	defer rawResp.Body.Close()

	err = json.NewDecoder(rawResp.Body).Decode(resp)
	if err != nil {
		return err
	}

	return nil
}

func (r *Realms) post(endpoint string, payload, resp interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	rawResp, err := r.c.Post(Domain+endpoint, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	err = json.NewDecoder(rawResp.Body).Decode(resp)
	if err != nil {
		return err
	}

	return nil
}
