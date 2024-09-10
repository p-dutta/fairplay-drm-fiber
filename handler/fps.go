package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fairplay-ksm/common"
	"fairplay-ksm/ksm"
	"fairplay-ksm/util"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type SpcMessage struct {
	ContentId    string `json:"contentId" validate:"required"`
	PackageId    string `json:"packageId" validate:"required"`
	ProviderId   string `json:"providerId" validate:"required"`
	DrmType      string `json:"drmType" validate:"required"`
	Payload      string `json:"payload" validate:"required"`
	SubscriberId string `json:"subscriberId" validate:"required"`
	DeviceType   string `json:"deviceType" validate:"required"`
	Token        string `json:"token" validate:"required"`
}

type CkcResult struct {
	Payload string `json:"payload" binding:"required"`
}

func GetLicense(ctx *fiber.Ctx) error {
	validate := validator.New()
	var spcMessage SpcMessage

	if err := ctx.BodyParser(&spcMessage); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(util.ErrorResponse("Invalid request body", err, 3001))
	}

	if err := validate.Struct(&spcMessage); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			util.ErrorResponse("Invalid request body", err, 3002))
	}

	if err := verifyToken(&spcMessage); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			util.ErrorResponse("You are not authorized to make this request", err, 3401))
	}

	fmt.Println("========================= Request ===============================")
	fmt.Printf("contentId - %v\n", spcMessage.ContentId)
	fmt.Printf("packageId - %v\n", spcMessage.PackageId)

	dataForKeyServer := common.DataForKeyServer{
		ContentId:  spcMessage.ContentId,
		PackageId:  spcMessage.PackageId,
		ProviderId: spcMessage.ProviderId,
		DrmType:    spcMessage.DrmType,
		Token:      spcMessage.Token,
	}

	k := &ksm.Ksm{
		Pub: common.PublicCert,
		Pri: common.PrivateKey,
		Rck: ksm.RandomContentKey{}, //NOTE: Don't use random key in your application.
		//Rck: "404df487de2e3d2b2d1993e86fabffa7", //NOTE: Don't use random key in your application.
		Ask: common.Ask,
	}

	var playback []byte
	var base64EncodingMethod string
	contentType := ctx.Get("Content-Type")

	if strings.Contains(spcMessage.Payload, "-") || strings.Contains(spcMessage.Payload, "_") {
		base64EncodingMethod = "URL"
		decoded, err := base64.URLEncoding.DecodeString(spcMessage.Payload)
		if err != nil {
			panic(err)
		}
		playback = decoded
	} else if strings.Contains(spcMessage.Payload, " ") && strings.Contains(spcMessage.Payload, "/") {
		base64EncodingMethod = "STD"
		decoded, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(spcMessage.Payload, " ", "+"))
		if err != nil {
			panic(err)
		}
		playback = decoded
	} else {
		base64EncodingMethod = "STD"
		decoded, err := base64.StdEncoding.DecodeString(spcMessage.Payload)
		if err != nil {
			panic(err)
		}
		playback = decoded
	}

	ckc, err := k.GenCKC(ctx, playback, dataForKeyServer)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			util.ErrorResponse("Could not generate CKC", err, 3003))
	}

	var result string

	switch base64EncodingMethod {
	case "URL":
		result = base64.URLEncoding.EncodeToString(ckc)
	case "STD":
		result = base64.StdEncoding.EncodeToString(ckc)
	default:
		result = base64.StdEncoding.EncodeToString(ckc)
	}

	// fmt.Println(result)

	switch contentType {
	case "application/json":
		//return ctx.Status(fiber.StatusOK).JSON(&CkcResult{Payload: result})
		return ctx.Status(fiber.StatusOK).JSON(util.SuccessResponse(&CkcResult{Payload: result}, "License data"))
	case "application/x-www-form-urlencoded":
		return ctx.Status(fiber.StatusOK).SendString("<payload>" + result + "</payload>")
	default:
		return ctx.Status(fiber.StatusOK).SendString("<payload>" + result + "</payload>")
	}

}

func verifyToken(spcMessage *SpcMessage) error {
	reqBody := common.VerifyTokenRequest{
		Token:        spcMessage.Token,
		SubscriberID: spcMessage.SubscriberId,
		ContentID:    spcMessage.ContentId,
		DeviceType:   spcMessage.DeviceType,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return errors.New("failed to encode request body")
	}
	verifyTokenApiURL := os.Getenv("VERIFY_TOKEN_API_URL")

	_, err = tokenVerificationServiceCall(verifyTokenApiURL, jsonBody)
	if err != nil {
		log.Println("Error while making request to key server:", err)
		return err
	}
	return nil
}

func tokenVerificationServiceCall(verifyTokenApiURL string, jsonBody []byte) ([]byte, error) {
	body, err := common.Cb.Execute(func() ([]byte, error) {
		req, err := http.NewRequest("POST", verifyTokenApiURL, bytes.NewBuffer(jsonBody))
		if err != nil {
			log.Println("Failed to create request for verifying token:", err)
			return nil, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API-Key", os.Getenv("VERIFY_TOKEN_X_API_KEY"))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error making POST request for verifying token:", err)
			return nil, err
		}
		defer resp.Body.Close()

		// Process the response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Failed to read response body for verifying token:", err)
			return nil, err
		}

		if resp.StatusCode == http.StatusOK {
			return body, nil
		} else {
			log.Printf("Verify Token API returned status code %d: %s", resp.StatusCode, string(body))
			return nil, errors.New("invalid token")
		}
	})
	if err != nil {
		log.Println("Error from cb execute:", err)
		return nil, err
	}

	return body, nil

}
