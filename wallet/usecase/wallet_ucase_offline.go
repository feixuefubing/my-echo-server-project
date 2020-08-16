package usecase

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	uuid "github.com/satori/go.uuid"
	"my-echo-server-project/domain"
	"strings"
)

type walletUsecaseOffline struct {
}

func (w walletUsecaseOffline) ExtractAddressFromSignedMessage(message string, signedMessage string) (address string, err error) {
	panic("离线版暂不实现")
}

func (w walletUsecaseOffline) SignMessage(message string, address string, password string) (signedMessage string, err error) {
	panic("离线版暂不实现")
}

func (w walletUsecaseOffline) NewAccount(password string) (address string, err error) {
	panic("离线版暂不实现")
}

func (w walletUsecaseOffline) GetAccounts() (accounts []string, err error) {
	panic("离线版暂不实现")
}

func NewWalletUsecaseOffline() domain.WalletUsecase {
	return &walletUsecaseOffline{}
}

func (w walletUsecaseOffline) GenerateRandomMessage() (message string, err error) {
	uuid := uuid.NewV4().String()
	return strings.ReplaceAll(uuid, "-", ""), nil
}

func (w walletUsecaseOffline) ValidatePrivateKeyAndAddress(message string, signedMessage string, expectedAddress string) (res bool, err error) {
	sig, err := hex.DecodeString(signedMessage[2:])
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	// see go-ethereum/internal/ethapi/api.go EcRecover
	if len(sig) != 65 || (sig[64] != 27 && sig[64] != 28) {
		return false, nil
	}
	sig[64] -= 27
	// TODO maybe last puzzle, why there is only one possible pubkey from elliptic curve?
	pubKey, err := crypto.SigToPub(signHash([]byte(message)), sig)
	if err != nil {
		return false,err
	}
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	return recoveredAddress == common.HexToAddress(expectedAddress), nil
}

// signHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calulcated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}