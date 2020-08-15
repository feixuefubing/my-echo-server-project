package dto
type VerifyReq struct {
	Address string `json:"address"`
	SignedMessage string `json:"signedMessage"`
}
