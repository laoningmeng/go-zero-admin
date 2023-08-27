package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
