package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_walletHttpHandler "my-echo-server-project/wallet/handler"
	_walletUcase "my-echo-server-project/wallet/usecase"
)

//var (
//	client *rpc.Client
//)

//func init() {
//	//获取geth rpc客户端
//	client, _ = rpc.Dial("http://localhost:8545")
//	if client == nil {
//		fmt.Println("rpc.Dial err")
//		panic("连接geth节点错误")
//		return
//	}
//}

func main() {
	// Echo实例
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	//}))

	// 实例化Handler层和业务层
	// 这里切换不同的实现
	//wu :=_walletUcase.NewWalletUsecase(client)
	wu :=_walletUcase.NewWalletUsecaseOffline()
	_walletHttpHandler.NewWalletHandler(e, wu)

	// 启动web服务
	e.Logger.Fatal(e.Start(":4000"))
	fmt.Print("echo server start")
}
