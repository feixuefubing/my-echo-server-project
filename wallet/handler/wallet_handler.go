package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"my-echo-server-project/domain"
	dto "my-echo-server-project/wallet/handler/dto"
	"net/http"
)

type WalletHandler struct {
	WUsecase domain.WalletUsecase
}

func NewWalletHandler(e *echo.Echo, us domain.WalletUsecase) {
	h := &WalletHandler{
		WUsecase: us,
	}
	//routes
	e.GET("/sign_in", h.SignIn)
	e.POST("/verify", h.Verify)

	e.GET("/getAccounts", h.GetAccounts)
	e.POST("/newAccount", h.NewAccount)
	e.POST("/signMessage", h.SignMessage)
	e.POST("/extractAddressFromSignedMessage", h.ExtractAddressFromSignedMessage)
}

func (h *WalletHandler) SignIn(c echo.Context) error {
	message, err := h.WUsecase.GenerateRandomMessage()
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	cookie := &http.Cookie{
		Name:     "signinmessage",
		Value:    message,
		HttpOnly: true,
	}
	c.SetCookie(cookie)
	resp := dto.SignInResp{Message: message}
	return c.JSON(http.StatusOK, resp)
}

func (h *WalletHandler) Verify(c echo.Context) error {
	// 取cookie中的随机message
	message, e := c.Cookie("signinmessage")
	if e != nil || message.Value == "" {
		fmt.Print(e.Error())
		return c.JSON(http.StatusOK, "请先登录")
	}
	//参数绑定 判空
	var req dto.VerifyReq
	c.Bind(&req)
	if req.SignedMessage == "" || req.Address == "" {
		return c.JSON(http.StatusOK, "参数为空")
	}
	isValid, err := h.WUsecase.ValidatePrivateKeyAndAddress(message.Value, req.SignedMessage, req.Address)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	resp := dto.VerifyResp{Verified: isValid}
	return c.JSON(http.StatusOK, resp)
}

//以下为测试geth功能
func (h *WalletHandler) GetAccounts(c echo.Context) error {
	accounts, err := h.WUsecase.GetAccounts()
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, accounts)
}

func (h *WalletHandler) NewAccount(c echo.Context) error {
	var req dto.NewAccountReq
	c.Bind(&req)
	address, err := h.WUsecase.NewAccount(req.Password)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, address)
}

func (h *WalletHandler) SignMessage(c echo.Context) error {
	var req dto.SignMessageReq
	c.Bind(&req)
	signedMessage, err := h.WUsecase.SignMessage(req.Message, req.Account, req.Password)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, signedMessage)
}

func (h *WalletHandler) ExtractAddressFromSignedMessage(c echo.Context) error {
	var req dto.ExtractMessageReq
	c.Bind(&req)
	address, err := h.WUsecase.ExtractAddressFromSignedMessage(req.Message, req.SignedMessage)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, address)
}
