package ksm

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fairplay-ksm/common"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"os"
)

// ContentKey is a interface that fetch asset content key and duration.
type ContentKey interface {
	FetchContentKey(ctx *fiber.Ctx, assetID []byte, dataForKeyServer common.DataForKeyServer) ([]byte, []byte, error)
	FetchContentKeyDuration(assetID []byte) (*CkcContentKeyDurationBlock, error)
}

var (
	_ ContentKey = RandomContentKey{}
)

// RandomContentKey is a object that implements ContentKey interface.
type RandomContentKey struct {
}

// FetchContentKey returns content key and iv for the given assetId.
func (RandomContentKey) FetchContentKey(ctx *fiber.Ctx, assetID []byte, dataForKeyServer common.DataForKeyServer) ([]byte, []byte, error) {
	/*key := make([]byte, 16)
	iv := make([]byte, 16)
	rand.Read(key)
	rand.Read(iv)
	*/

	/*staticKey := os.Getenv("STATIC_KEY_HEX")
	staticIv := os.Getenv("STATIC_IV_HEX")
	key, _ := hex.DecodeString(staticKey)
	iv, _ := hex.DecodeString(staticIv)*/
	var keyServerKey []byte
	var keyServerIv []byte

	keyServerBaseUrl := os.Getenv("KEY_SERVER_BASE_URL")
	contentId := dataForKeyServer.ContentId
	packageId := dataForKeyServer.PackageId

	keyServerUri := keyServerBaseUrl + contentId + "/" + packageId

	requestBody := common.KeyServerRequest{
		DrmScheme: []string{common.DrmFP},
		Quality:   []string{common.QualityHD, common.QualityAudio, common.QualitySD},
	}

	// Convert struct to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Println("Error marshalling request body:", err)
		return nil, nil, err
	}

	//return body.([]byte), nil

	// Example request to the key server
	/*resp, err := http.Post(keyServerUri, "application/json", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		log.Println("Error making POST request to key server:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()
	response, err := getDataFromKeyServer(keyServerUri, requestBodyJSON)
	err = json.NewDecoder(resp.Body).Decode(&keyServerResponse)
	if err != nil {
		log.Println("Error decoding key server response:", err)
		return nil, nil, err
	}

	*/

	response, err := getDataFromKeyServer(keyServerUri, requestBodyJSON)
	if err != nil {
		log.Println("Error while making request to key server:", err)
		return nil, nil, err
	}

	var keyServerResponse common.KeyServerResponse

	//Decode the response

	//json.Unmarshal doesn't work here. You'll have to do:
	/* response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, nil, err
	}*/
	//and then pass body as : json.Unmarshal(response, &keyServerResponse)

	if err := json.Unmarshal(response, &keyServerResponse); err != nil {
		log.Println("Error decoding key server response:", err)
		return nil, nil, err
	}

	// Example of handling the response
	if keyServerResponse.Success {
		for _, keyMap := range keyServerResponse.Data.Keys {
			for _, key := range keyMap {
				fmt.Println("Key ID:", key.KeyID)
				fmt.Println("Key IV:", key.KeyIV)
				keyServerKey, _ = hex.DecodeString(key.Key)
				keyServerIv, _ = hex.DecodeString(key.KeyIV)
				break
			}
			break
		}
	} else {
		return nil, nil, errors.New(keyServerResponse.Message)
	}

	return keyServerKey, keyServerIv, nil
}

// FetchContentKeyDuration returns CkcContentKeyDurationBlock for the given assetId.
func (RandomContentKey) FetchContentKeyDuration(assetID []byte) (*CkcContentKeyDurationBlock, error) {

	// LeaseDuration := mathRand.Uint32()  // The duration of the lease, if any, in seconds.
	// RentalDuration := mathRand.Uint32() // The duration of the rental, if any, in seconds.

	LeaseDuration := uint32(3600)
	RentalDuration := uint32(3600)

	return NewCkcContentKeyDurationBlock(LeaseDuration, RentalDuration), nil
}

// CKCPayload is a object that store ckc payload.
type CKCPayload struct {
	SK             []byte //Session key
	HU             []byte
	R1             []byte
	IntegrityBytes []byte
}

func getDataFromKeyServer(keyServerUri string, requestBodyJSON []byte) ([]byte, error) {

	//cb := config.CircuitBreakerConfig()

	body, err := common.Cb.Execute(func() ([]byte, error) {
		resp, err := http.Post(keyServerUri, "application/json", bytes.NewBuffer(requestBodyJSON))
		if err != nil {
			log.Println("Error making POST request to key server:", err)
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	})
	if err != nil {
		log.Println("Error from cb execute:", err)
		return nil, err
	}

	return body, nil

}
