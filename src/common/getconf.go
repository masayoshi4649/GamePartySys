package common

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/BurntSushi/toml"
)

// ConfigToml ... config.toml
type ConfigToml struct {
	App      App
	Discord  Discord
	Postgres Postgres
}

// App ... [App]
type App struct {
	Myhostname string
	Port       string
	Session    string
}

// Discord ... [Discord]
type Discord struct {
	ClientID     string
	ClientSecret string
	AuthURI      string
	AuthScope    string
	OauthDir     string
	LoginURL     string
	TokenURI     string
}

// Postgres ... [Postgres]
type Postgres struct {
	DatabaseName string
	Host         string
	Port         string
	PsqlID       string
	PsqlPass     string
	ConnectStr   string
}

// GetConf ... Read config
func getConf() ConfigToml {
	var conf ConfigToml
	_, err := toml.DecodeFile("config.toml", &conf)
	if err != nil {
		panic(err)
	}

	// 接続文字列
	conf.Postgres.ConnectStr =
		"host=" + conf.Postgres.Host + " " +
			"port=" + conf.Postgres.Port + " " +
			"user=" + conf.Postgres.PsqlID + " " +
			"password=" + conf.Postgres.PsqlPass + " " +
			"dbname=" + conf.Postgres.DatabaseName + " " +
			"sslmode=" + "disable"

	// 認証用URL
	redirectURI := conf.App.Myhostname + conf.Discord.OauthDir

	// 【clientID】
	discordURI := strings.Replace(conf.Discord.AuthURI, "【clientID】", conf.Discord.ClientID, -1)
	// 【enc_redirect_uri】
	discordURI = strings.Replace(discordURI, "【enc_redirect_uri】", url.QueryEscape(redirectURI), -1)
	// 【enc_scope】
	discordURI = strings.Replace(discordURI, "【enc_scope】", conf.Discord.AuthScope, -1)
	conf.Discord.LoginURL = discordURI
	fmt.Println("discordURI=>", discordURI)

	return conf
}

// Config ... メモリ上に配置(アプリ再起動まで再読み込みしないため)
var Config ConfigToml

func init() {
	Config = getConf()
}
