package dto

type PostLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostCheckTokenReq struct {
	Token string `json:"token"`
}
