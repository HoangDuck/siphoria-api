package services

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"hotel-booking-api/helper"
	"hotel-booking-api/model"
	"io"
	"io/ioutil"
	"net/http"
)

type FacebookOauthService struct {
	FacebookLoginConfig oauth2.Config
}

var oauthFacebookService *FacebookOauthService

const OauthFacebookUrlAPI = "https://graph.facebook.com/me?access_token="

func GetFacebookOauth2ServiceInstance() *FacebookOauthService {
	if oauthFacebookService == nil {
		oauthFacebookService = new(FacebookOauthService)
		oauthFacebookService.SetUpConfig(*ConfigInfo)
	}
	return oauthFacebookService
}

func (oauth *FacebookOauthService) SetUpConfig(cfg model.Config) {
	// Oauth configuration for Google
	oauth.FacebookLoginConfig = oauth2.Config{
		ClientID:     cfg.FacebookOauth2.ClientId,
		ClientSecret: cfg.FacebookOauth2.ClientSec,
		Endpoint:     facebook.Endpoint,
		RedirectURL:  cfg.FacebookOauth2.RedirectUrl,
		Scopes:       []string{},
	}
	if oauth.FacebookLoginConfig.ClientID != "" {
		fmt.Printf("Connected Oauth2")
	} else {
		fmt.Printf("Connect to Oauth2 failed")
	}
}

func (oauth *FacebookOauthService) FacebookAuthenticationService(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Create oauthState cookie
	oauthState := helper.GenerateStateOauthCookie(w)
	u := oauth.FacebookLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (oauth *FacebookOauthService) FacebookAuthenticationCallBack(w http.ResponseWriter, r *http.Request) map[string]interface{} {
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
		_, err := fmt.Fprintf(w, "invalid oauth facebook state")
		if err != nil {
			return nil
		}
		return nil
	}

	// Exchange Auth Code for Tokens
	token, err := oauth.FacebookLoginConfig.Exchange(
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
	response, err := http.Get(OauthFacebookUrlAPI + token.AccessToken)

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
