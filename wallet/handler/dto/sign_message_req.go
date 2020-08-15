package dto
type SignMessageReq struct {
	Message string `json:"message"`
	Account string `json:"address"`
	Password string `json:"password"`
}
