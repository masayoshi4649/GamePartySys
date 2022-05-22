package discord

// Discordtoken ... {"access_token": "_____", "expires_in": 604800, "refresh_token": "_____", "scope": "identify email", "token_type": "Bearer"}
type Discordtoken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// Me ... {"id": "_____", "username": "_____", "avatar": "_____", "discriminator": "____", "public_flags": 0, "flags": 0, "email": "_____", "verified": true, "locale": "ja", "mfa_enabled": false}
type Me struct {
	ID            string `json:"id"`
	UserName      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
	Flags         int    `json:"flags"`
	Email         string `json:"email"`
	Verified      bool   `json:"verified"`
	Locale        string `json:"locale"`
	MfaEnabled    bool   `json:"mfa_enabled"`
}
