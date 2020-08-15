package usecase

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	bytes, err := hex.DecodeString(signedMessage[2:])
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	byteAddress, err := EcRecover([]byte(message), bytes)
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	address := byteAddress.String()
	return strings.ToLower(address) == expectedAddress, nil
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

//// Sign calculates an Ethereum ECDSA signature for:
//// keccack256("\x19Ethereum Signed Message:\n" + len(message) + message))
////
//// Note, the produced signature conforms to the secp256k1 curve R, S and V values,
//// where the V value will be 27 or 28 for legacy reasons.
////
//// The key used to calculate the signature is decrypted with the given password.
////
//// https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal_sign
//func (s *PrivateAccountAPI) Sign(ctx context.Context, data hexutil.Bytes, addr common.Address, passwd string) (hexutil.Bytes, error) {
//	// Look up the wallet containing the requested signer
//	account := accounts.Account{Address: addr}
//
//	wallet, err := s.b.AccountManager().Find(account)
//	if err != nil {
//		return nil, err
//	}
//	// Assemble sign the data with the wallet
//	signature, err := wallet.SignHashWithPassphrase(account, passwd, signHash(data))
//	if err != nil {
//		log.Warn("Failed data sign attempt", "address", addr, "err", err)
//		return nil, err
//	}
//	signature[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
//	return signature, nil
//}

// EcRecover returns the address for the account that was used to create the signature.
// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
// the address of:
// hash = keccak256("\x19Ethereum Signed Message:\n"${message length}${message})
// addr = ecrecover(hash, signature)
//
// Note, the signature must conform to the secp256k1 curve R, S and V values, where
// the V value must be 27 or 28 for legacy reasons.
//
// https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal_ecRecover
func  EcRecover( data, sig hexutil.Bytes) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1

	rpk, err := crypto.SigToPub(signHash(data), sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*rpk), nil
}
