package dto
type ExtractMessageReq struct {
	Message string `json:"message"`
	SignedMessage string `json:"signedMessage"`
}
