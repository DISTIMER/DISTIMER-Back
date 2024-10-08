package crypt

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"pentag.kr/distimer/configs"
)

type GooglePayload struct {
	SUB           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

func convertToken(accessToken string) (*GooglePayload, error) {
	// call http request to this URL, this is a valid
	// URL provided from google to convert access token into
	// valid user data
	resp, err := resty.New().R().Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
	if err != nil {
		return nil, err
	} else if resp.StatusCode() != 200 {
		return nil, errors.New("invalid token")
	}
	// Unmarshal raw response body to a map
	var body map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &body); err != nil {
		return nil, err
	}

	// If json body containing error,
	// then the token is indeed invalid. return invalid token err
	if body["error"] != nil {
		return nil, errors.New("invalid token")
	}

	// Bind JSON into struct
	var data GooglePayload
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type GoogleTokenPayload struct {
	AccessToken string `json:"access_token"`
}

func VerifyGoogleToken(code string) (*GooglePayload, error) {
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"code":          code,
			"client_id":     configs.Env.GoogleClientID,
			"client_secret": configs.Env.GoogleClientSecret,
			"grant_type":    "authorization_code",
		}).
		Post("https://oauth2.googleapis.com/token")
	if err != nil {
		return nil, err
	}
	payload := GoogleTokenPayload{}
	err = json.Unmarshal(resp.Body(), &payload)
	if err != nil {
		return nil, err
	}
	profile, err := convertToken(payload.AccessToken)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
