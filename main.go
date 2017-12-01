package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	rand.Seed(time.Now().Unix())

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

const spotifyClientID = "7b55bd86c0484c67b75e1cae230781be"
const spotifyClientSecret = "OH YOU DUMMY"
const spotifyRedirectURI = "http://localhost:8080/callback"

func spotifyHandler(w http.ResponseWriter, r *http.Request) {
	state := &http.Cookie{
		Name:  "spotify_auth_state",
		Value: randSeq(16),
	}

	http.SetCookie(w, state)

	v := url.Values{}
	v.Set("response_type", "code")
	v.Set("client_id", spotifyClientID)
	v.Set("scope", "user-read-private user-read-email")
	v.Set("redirect_uri", spotifyRedirectURI)
	v.Set("state", state.String())
	http.Redirect(w, r, "https://accounts.spotify.com/authorize?"+v.Encode(), http.StatusFound)
}

func spotifyCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func main() {
	http.HandleFunc("/spotify", spotifyHandler)
	http.HandleFunc("/callback", spotifyCallbackHandler)
	http.ListenAndServe(":8080", nil)
}
