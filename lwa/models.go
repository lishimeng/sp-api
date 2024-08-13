package lwa

import "time"

const expireBefore = 30

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"` //seconds
	RefreshToken string `json:"refresh_token"`
	ApplyAt      int64  `json:"apply_at,omitempty"` // 申请时间戳
}

// RefreshTime 记录申请时间
func (at *AccessToken) RefreshTime() {
	at.ApplyAt = time.Now().Unix()
}

func (at *AccessToken) Valid() (expired bool) {

	if at.ExpiresIn > 0 {
		expired = time.Now().Unix() < at.ApplyAt+int64(at.ExpiresIn)+expireBefore
	}
	return
}
