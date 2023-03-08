package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/validator.v2"
)

type Handle struct {
	service Servicer
}

func NewHandle(service Servicer) *Handle {
	return &Handle{service: service}
}

func (h *Handle) ExecuteGetAllWallet(c echo.Context) error {
	// fmt.Println("hi test")
	data := h.service.ApiGetAllWallet(c)

	return c.JSON(http.StatusOK, data)
}

func (h *Handle) ExecuteCreateNewWallet(c echo.Context) error {
	var req RequestNewWallet
	var res ResponseNewWallet

	if err := c.Bind(&req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := validator.Validate(req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	data := h.service.ApiPostCreatNewWallet(c, req)

	return c.JSON(http.StatusOK, data)
}

func (h *Handle) ExecuteGetWalletDetail(c echo.Context) error {
	var req RequestWalletDetail
	var res ResponseWalletDetail

	if err := c.Bind(&req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := validator.Validate(req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	data := h.service.ApiGetWalletDetail(c, req)

	return c.JSON(http.StatusOK, data)
}

func (h *Handle) ExecuteAddBalance(c echo.Context) error {
	var req RequestAddBalance
	var res ResponseAddBalance

	if err := c.Bind(&req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := validator.Validate(req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	data := h.service.ApiPostAddBalance(c, req)

	return c.JSON(http.StatusOK, data)
}

func (h *Handle) ExecuteDeductBalance(c echo.Context) error {
	var req RequestDeductBalance
	var res ResponseDeductBalance

	if err := c.Bind(&req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := validator.Validate(req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	data := h.service.ApiPostDeductBalance(c, req)

	return c.JSON(http.StatusOK, data)
}

func (h *Handle) ExecuteWalletStatus(c echo.Context) error {
	var req RequestWalletStatus
	var res ResponseWalletStatus

	if err := c.Bind(&req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := validator.Validate(req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	data := h.service.ApiPostWalletStatus(c, req)

	return c.JSON(http.StatusOK, data)
}

func (h *Handle) ExecuteTest1(c echo.Context) error {
	var req ReqTest1
	var res ResTest1
	var reqH ReqHeaderTest1
	// var header http.Header
	// reqH.AppID = header.Get("app-id")

	SetHeader(c.Request().Header, &reqH)

	// fmt.Printf("req =%v\n", req)
	if err := c.Bind(&req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	if err := validator.Validate(req); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}
	// fmt.Printf("header req =%v\n", reqH)
	if err := validator.Validate(reqH); err != nil {
		res.RespCode = "SV002"
		res.RespDesc = "invalid request header"
		res.RespDisplay = err.Error()
		return c.JSON(http.StatusBadRequest, res)
	}

	data := h.service.ApiTest1(c, req)

	return c.JSON(http.StatusOK, data)
}

func SetHeader(header http.Header, reqHeader *ReqHeaderTest1) {
	reqHeader.AppID = header.Get("app-id")

}
