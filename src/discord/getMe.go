package discord

import (
	"encoding/json"
	"net/http"
)

func getMe(accessToken string, tokenType string) (userdata Me) {
	// make HTTP Client
	client := &http.Client{}

	// make Request
	req, _ := http.NewRequest("GET", "https://discordapp.com/api/users/@me", nil)
	header := http.Header{}
	var authorizationStr string = tokenType + " " + accessToken
	header.Add("Authorization", authorizationStr)
	req.Header = header

	// do Request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var me Me
	err = json.NewDecoder(resp.Body).Decode(&me)

	return me
}
