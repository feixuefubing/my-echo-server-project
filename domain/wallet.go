package domain

type Wallet struct {
	Password string
	Address  string
}

type WalletUsecase interface {
	GenerateRandomMessage() (message string, err error)
	ValidatePrivateKeyAndAddress(message string, signedMessage string, expectedAddress string) (res bool, err error)
	//以下为测试geth功能
	ExtractAddressFromSignedMessage(message string, signedMessage string) (address string, err error)
	SignMessage(message string, address string, password string) (signedMessage string, err error)
	NewAccount(password string) (address string, err error)
	GetAccounts() (accounts []string, err error)
}
