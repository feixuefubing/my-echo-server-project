package usecase_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"my-echo-server-project/domain"
	"my-echo-server-project/wallet/usecase"
	"my-echo-server-project/wallet/util"
	"os"
	"testing"
)

var (
	ucase domain.WalletUsecase
)

func TestMain(m *testing.M) {
	//获取geth rpc客户端
	client, _ := rpc.Dial("http://localhost:8545")
	if client == nil {
		fmt.Println("rpc.Dial err")
		panic("连接geth节点错误")
		return
	}
	ucase = usecase.NewWalletUsecase(client)
	os.Exit(m.Run())
}

func TestGenerateRandomMessage(t *testing.T) {
	//50个统计上认为大样本
	var messageArr [50]string
	messageArr[0], _ = ucase.GenerateRandomMessage()
	for i := 1; i < 50; i++ {
		messageArr[i], _ = ucase.GenerateRandomMessage()
		_, ok := interface{}(messageArr[i]).(string)
		//为字符串类型且长度固定为32
		if !ok || len(messageArr[i]) != 32 {
			t.Fatalf("random message generate wrong")
		}
	}
	//50个随机message不重复
	if util.ContainsDuplicate(messageArr[:]) {
		t.Fatalf("random message generate duplicate")
	}
}

func TestValidatePrivateKeyAndAddress(t *testing.T) {
	cases := []struct {
		message          string
		signedMessage    string
		expectedAddress  string
		expectedValidRes bool
	}{
		//正确用例1 signed by different geth node
		{"aer3r", "0xac7cb5757ef47acbd00d6395c509f746dc3beea0100407ac5e1a8cf5405d5a4c17dddc4b627ad7968bbb3da3a107a7474e707bf959b909293ea5e8bf361bc8951b", "0x2b3a8e605ccffd20b38e357281dde106e7b9d96f", true},
		//正确用例2
		{"864d3d085e2d45d08fa3652abc67050e", "0x11deebe9ad4063e13ef5e51ffcb35c215913830e909ffd60aef0d45753c2e9163cfe4152387893c5d45f03e2d09d56c003893f2a2d462d65855c3c5a010f72241b", "0x887d8b0352d5ed525608edf5b5b2fc8e15fc52db", true},
		//message错误
		{"864d3d085e2d45d08fa3652abc67050e", "0x11deebe9ad4063e13ef5e51ffcb35c215913830e909ffd60aef0d45753c2e9163cfe4152387893c5d45f03e2d09d56c003893f2a2d462d65855c3c5a010f72241b", "0x887d8b0352d5ed525608edf5b5b2fc8e15fc52dc", false},
		//expectedAddress错误
		{"864d3d085e2d45d08fa3652abc67050a", "0x11deebe9ad4063e13ef5e51ffcb35c215913830e909ffd60aef0d45753c2e9163cfe4152387893c5d45f03e2d09d56c003893f2a2d462d65855c3c5a010f72241b", "0x887d8b0352d5ed525608edf5b5b2fc8e15fc52db", false},
		//signedMessage错误
		{"864d3d085e2d45d08fa3652abc67050e", "0x11deebe9ad4063e13e", "0x887d8b0352d5ed525608edf5b5b2fc8e15fc52db", false},
	}

	for _, c := range cases {
		isValid, _ := ucase.ValidatePrivateKeyAndAddress(c.message, c.signedMessage, c.expectedAddress)
		if isValid != c.expectedValidRes {
			t.Fatalf("ValidatePrivateKeyAndAddress failed, message: %s, signedMessage:%s, expectedAddress:%s, isValid:%t", c.message, c.signedMessage, c.expectedAddress, isValid)
		}
	}
}
