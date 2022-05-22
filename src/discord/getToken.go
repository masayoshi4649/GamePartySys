package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/masayoshi4649/GamePartySys/src/common"
	"github.com/masayoshi4649/GamePartySys/src/dbpsql"
)

// GetToken ... 認証codeを変換する
func GetToken(code string) (userid string) {
	conf := common.Config
	redirectURI := conf.App.Myhostname + conf.Discord.OauthDir

	// make HTTP Client
	client := &http.Client{}

	// make Request
	var datastr string = "client_id=" + conf.Discord.ClientID + "&client_secret=" + conf.Discord.ClientSecret + "&grant_type=authorization_code&code=" + code + "&redirect_uri=" + redirectURI + "&scope=" + conf.Discord.AuthScope

	req, _ := http.NewRequest("POST", conf.Discord.TokenURI, bytes.NewBufferString(datastr))
	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header = header

	// do Request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var discordtoken Discordtoken
	err = json.NewDecoder(resp.Body).Decode(&discordtoken)

	userdata := getMe(discordtoken.AccessToken, discordtoken.TokenType)

	fmt.Println("userdata.ID=>", userdata.ID)
	if userdata.ID != "" {
		upsertDiscordToken(userdata.ID, discordtoken)
		upsertDiscordInfo(userdata)
	}

	return userdata.ID
}

// upsertDiscordToken ... アクセストークンとリフレッシュトークンを保存
func upsertDiscordToken(id string, token Discordtoken) {
	dbpsql.UpsertDToken(id, token.AccessToken, token.ExpiresIn, token.RefreshToken, token.Scope, token.TokenType)
}

// upsertDiscordInfo ... Discordの情報を保存
func upsertDiscordInfo(me Me) {
	dbpsql.UpsertDInfo(me.ID, me.UserName, me.Avatar, me.Discriminator, me.PublicFlags, me.Flags, me.Email, me.Verified, me.Locale, me.MfaEnabled)
}
