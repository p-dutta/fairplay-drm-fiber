package common

import (
	"crypto/rsa"
	"fairplay-ksm/config"
)

var PrivateKey *rsa.PrivateKey
var PublicCert *rsa.PublicKey
var Ask []byte //Application Secret Key

// var Cb *gobreaker.CircuitBreaker[[]byte]

var Cb = config.CircuitBreakerConfig()

const (
	DrmWV = "WV"
	DrmFP = "FP"
	DrmPR = "PR"
)

const (
	QualityAudio = "AUDIO"
	QualitySD    = "SD"
	QualityHD    = "HD"
	QualityUHD1  = "UHD1"
	QualityUHD2  = "UHD2"
)

type DataForKeyServer struct {
	ContentId  string `json:"contentId" validate:"required"`
	PackageId  string `json:"packageId" validate:"required"`
	ProviderId string `json:"providerId" validate:"required"`
	DrmType    string `json:"drmType" validate:"required"`
	Token      string `json:"token" validate:"required"`
}

type KeyServerRequest struct {
	DrmScheme []string `json:"drmScheme"`
	Quality   []string `json:"quality" validate:"required"`
}

type KeyServerResponse struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Data struct {
	ContentID  string           `json:"contentId"`
	PackageID  string           `json:"packageId"`
	ProviderID string           `json:"providerId"`
	Keys       []map[string]Key `json:"keys"`
}

type Key struct {
	KeyID string `json:"keyId"`
	KeyIV string `json:"keyIv"`
	Key   string `json:"key"`
}

type VerifyTokenRequest struct {
	Token        string `json:"token"`
	SubscriberID string `json:"subscriberId"`
	ContentID    string `json:"contentId"`
	DeviceType   string `json:"deviceType"`
}

type VerifyTokenResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
