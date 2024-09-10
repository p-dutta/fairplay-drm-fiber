// NOTE: In production, please implement your content key logic instead of Random Content Key
package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"

	"fairplay-ksm/ksm"
)

type RandomContentKey struct {
}

// FetchContentKey Implement FetchContentKey func
func (RandomContentKey) FetchContentKey(assetID []byte) ([]byte, []byte, error) {
	return generateDummyKeyIVPair(assetID)
}

// FetchContentKeyDuration Implement FetchContentKeyDuration func
func (RandomContentKey) FetchContentKeyDuration(assetID []byte) (*ksm.CkcContentKeyDurationBlock, error) {

	LeaseDuration := rand.Uint32()  // The duration of the lease, if any, in seconds.
	RentalDuration := rand.Uint32() // The duration of the rental, if any, in seconds.

	return ksm.NewCkcContentKeyDurationBlock(LeaseDuration, RentalDuration), nil
}

func generateDummyKeyIVPair(assetID []byte) ([]byte, []byte, error) {
	/*dummyKey := make([]byte, 16)
	dummyIV := make([]byte, 16)
	rand.Read(dummyIV)*/
	staticKey := os.Getenv("STATIC_KEY_HEX")
	staticIv := os.Getenv("STATIC_IV_HEX")
	dummyKey, _ := hex.DecodeString(staticKey)
	dummyIV, _ := hex.DecodeString(staticIv)

	if len(assetID) == 0 {
		rand.Read(dummyKey)
		return dummyKey, dummyIV, nil
	}
	// NOTE: Here is to implement your key generator.
	generator := md5.New()
	generator.Write(assetID)
	dummyKey = generator.Sum(nil)
	return dummyKey, dummyIV, nil
}
