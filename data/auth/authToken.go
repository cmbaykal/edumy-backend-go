package authToken

type AuthToken struct {
	AuthToken  string `json:"token"`
	ExpireTime string `json:"expireTime"`
}
