package usecase

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"my-echo-server-project/domain"
	"my-echo-server-project/wallet/util"
	"strings"
)

type walletUsecase struct {
	client *rpc.Client
}

func NewWalletUsecase(rpcClient *rpc.Client) domain.WalletUsecase {
	return &walletUsecase{client: rpcClient}
}

func (w walletUsecase) GenerateRandomMessage() (message string, err error) {
	uuid := uuid.NewV4().String()
	return strings.ReplaceAll(uuid, "-", ""), nil
}

func (w walletUsecase) ValidatePrivateKeyAndAddress(message string, signedMessage string, expectedAddress string) (res bool, err error) {
	var address string
	err = w.client.Call(&address, "personal_ecRecover", util.TransferToHexStr(message), signedMessage)
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	return address == expectedAddress, nil
}

func (w walletUsecase) ExtractAddressFromSignedMessage(message string, signedMessage string) (address string, err error) {
	err = w.client.Call(&address, "personal_ecRecover", util.TransferToHexStr(message), signedMessage)
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}
	return address, nil
}

func (w walletUsecase) SignMessage(message string, address string, password string) (signedMessage string, err error) {
	err = w.client.Call(&signedMessage, "personal_sign", util.TransferToHexStr(message), address, password)
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}
	return signedMessage, nil
}

func (w walletUsecase) NewAccount(password string) (address string, err error) {
	err = w.client.Call(&address, "personal_newAccount", password)
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}
	return address, nil
}

func (w walletUsecase) GetAccounts() (accounts []string, err error) {
	err = w.client.Call(&accounts, "eth_accounts")
	if err != nil {
		fmt.Print(err.Error())
		return nil, errors.New("账户列表获取错误")
	}
	return accounts, nil
}
