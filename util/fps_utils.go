package util

import (
	"encoding/base64"
	"encoding/hex"
	"fairplay-ksm/common"
	"fairplay-ksm/cryptos"
	"fmt"
	"os"
)

func envBase64Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("Env Decode Error: ", err)
		return []byte{}
	}
	return data
}

func ReadPublicCert() {
	pubEnvVar := envBase64Decode(os.Getenv("FAIRPLAY_CERTIFICATION"))

	if len(pubEnvVar) == 0 {
		panic("Can't not find FAIRPLAY CERTIFICATION")
	}

	pubCert, err := cryptos.ParsePublicCertification(pubEnvVar)
	if err != nil {
		panic(err)
	}
	common.PublicCert = pubCert
}

func ReadPriKey() {
	priEnvVar := envBase64Decode(os.Getenv("FAIRPLAY_PRIVATE_KEY"))

	if len(priEnvVar) == 0 {
		panic("Can't not find FAIRPLAY PRIVATE KEY")
	}
	priKey, err := cryptos.DecryptPriKey(priEnvVar, []byte(""))
	if err != nil {
		panic(err)
	}
	common.PrivateKey = priKey
}

func ReadASk() {
	askEnvVar := os.Getenv("FAIRPLAY_APPLICATION_SERVICE_KEY")
	if len(askEnvVar) == 0 {
		askEnvVar = "d87ce7a26081de2e8eb8acef3a6dc179" //Apple provided
	}
	ask, err := hex.DecodeString(askEnvVar)
	if err != nil {
		panic(err)
	}
	common.Ask = ask
}
