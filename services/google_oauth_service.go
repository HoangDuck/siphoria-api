package services

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"hotel-booking-api/helper"
	"hotel-booking-api/model"
	"io"
	"io/ioutil"
	"net/http"
)

type GoogleOauthService struct {
	GoogleLoginConfig oauth2.Config
}

var oauthService *GoogleOauthService

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?id_token="

func GetOauth2ServiceInstance() *GoogleOauthService {
	if oauthService == nil {
		oauthService = new(GoogleOauthService)
		oauthService.SetUpConfig(*ConfigInfo)
	}
	return oauthService
}

func (oauth *GoogleOauthService) SetUpConfig(cfg model.Config) {
	// Oauth configuration for Google
	oauth.GoogleLoginConfig = oauth2.Config{
		ClientID:     cfg.Oauth2.ClientId,
		ClientSecret: cfg.Oauth2.ClientSec,
		Endpoint:     google.Endpoint,
		RedirectURL:  cfg.Oauth2.RedirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
	if oauth.GoogleLoginConfig.ClientID != "" {
		fmt.Printf("Connected Oauth2")
	} else {
		fmt.Printf("Connect to Oauth2 failed")
	}
}

func (oauth *GoogleOauthService) GoogleAuthenticationService(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Create oauthState cookie
	oauthState := helper.GenerateStateOauthCookie(w)
	u := oauth.GoogleLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (oauth *GoogleOauthService) GetUserInfoWithToken(token string) map[string]interface{} {
	response, err := http.Get(OauthGoogleUrlAPI + token)

	// ERROR : Unable to get user data from google
	if err != nil {
		return nil
	}

	// Parse user data JSON Object
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	value := map[string]interface{}{}
	if err = json.Unmarshal(contents, &value); err != nil {
		panic(err)
	}
	// send back response to browser
	return value
}

func (oauth *GoogleOauthService) AuthenticationCallBack(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// check is method is correct
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	// get oauth state from cookie for this user
	oauthState, _ := r.Cookie("oauthstate")
	state := r.FormValue("state")
	code := r.FormValue("code")
	w.Header().Add("content-type", "application/json")
	if state != oauthState.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		_, err := fmt.Fprintf(w, "invalid oauth google state")
		if err != nil {
			return nil
		}
		return nil
	}

	// Exchange Auth Code for Tokens
	token, err := oauth.GoogleLoginConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		_, err := fmt.Fprintf(w, "falied code exchange: %s", err.Error())
		if err != nil {
			return nil
		}
		return nil
	}

	// Fetch User Data from google server
	response, err := http.Get(OauthGoogleUrlAPI + token.AccessToken)

	// ERROR : Unable to get user data from google
	if err != nil {
		_, err := fmt.Fprintf(w, "failed getting user info: %s", err.Error())
		if err != nil {
			return nil
		}
		return nil
	}

	// Parse user data JSON Object
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		_, err := fmt.Fprintf(w, "failed read response: %s", err.Error())
		if err != nil {
			return nil
		}
		return nil
	}
	value := map[string]interface{}{}
	if err = json.Unmarshal(contents, &value); err != nil {
		panic(err)
	}
	// send back response to browser
	return value
}
