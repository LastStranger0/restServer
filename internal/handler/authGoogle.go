package handler

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"restserver/internal/tokens"
)

var (
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://127.0.0.1:8080/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	randomState = "random"
)

// SignIn ...
func SignIn(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Host+"/login", http.StatusTemporaryRedirect)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallBack(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Println("wrong!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("wrong!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	var myID = idStruct{}
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &myID)

	tokenStr, err := tokens.GenerateToken(myID.Id)
	if err != nil {
		log.Println("err", err)
		http.Error(w, "ooops", http.StatusInternalServerError)
		return
	}

	c := &http.Cookie{Name: "X-Session-Token", Value: tokenStr}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

type idStruct struct {
	Id string `json:"id"`
}
