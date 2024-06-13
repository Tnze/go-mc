package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func usernameToUUID(name string) (string, uuid.UUID, error) {
	var id uuid.UUID
	resp, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)
	if err != nil {
		return "", id, err
	}

	var body struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return "", id, err
	}

	id, err = uuid.Parse(body.ID)
	return body.Name, id, err
}

func intArrayToUUID(uuidInts []int32) (id uuid.UUID, err error) {
	if uuidLen := len(uuidInts); uuidLen != 4 {
		err = fmt.Errorf("invalid UUID len: %d * int32", uuidLen)
		return
	}
	for i, v := range uuidInts {
		binary.BigEndian.PutUint32(id[i*4:], uint32(v))
	}
	return
}

func uuidToIntArray(id uuid.UUID) (ints [4]int32) {
	for i := range ints {
		ints[i] = int32(binary.BigEndian.Uint32(id[i*4:]))
	}
	return
}
