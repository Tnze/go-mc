package userApi

import (
	"time"
)

type KeyPairResp struct {
	KeyPair struct {
		PrivateKey string `json:"privateKey"`
		PublicKey  string `json:"publicKey"`
	} `json:"keyPair"`
	PublicKeySignature string    `json:"publicKeySignature"`
	ExpiresAt          time.Time `json:"expiresAt"`
	RefreshedAfter     time.Time `json:"refreshedAfter"`
}

func GetOrFetchKeyPair(accessToken string) (KeyPairResp, error) {
	return fetchKeyPair(accessToken) // TODO: cache
}

func fetchKeyPair(accessToken string) (KeyPairResp, error) {
	var keyPairResp KeyPairResp
	err := post("/player/certificates", accessToken, &keyPairResp)
	return keyPairResp, err
}
