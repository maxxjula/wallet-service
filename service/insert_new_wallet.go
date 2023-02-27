package service

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/maxxjula/wallet-service/repository"
)

func (s *Service) ApiPostCreatNewWallet(c echo.Context, req RequestNewWallet) *ResponseNewWallet {

	var res ResponseNewWallet

	dataWallet := repository.TbWallet{
		CitizenId:   req.CitizenId,
		Title:       req.Title,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		Address:     req.Address,
		Subdistrict: req.Subdistrict,
		District:    req.District,
		Province:    req.Province,
		ZipCode:     req.ZipCode,
		Tel:         req.Tel,
		CreatedDt:   time.Now(),
		CreatedBy:   "admin"}
	dataWalletAccount := repository.TbWalletAccount{
		CitizenId: req.CitizenId,
		AccountNo: req.AccountNo,
		Balance:   req.Amount,
		Status:    "A",
		Remark:    req.Remark,
		CreatedDt: time.Now(),
		CreatedBy: "admin"}

	dataTrans := repository.TbTransaction{
		Type:      "D",
		AccountNo: req.AccountNo,
		Amount:    req.Amount,
		TransDt:   time.Now(),
		CreatedDt: time.Now(),
		CreatedBy: "admin"}

	result, err := s.repository.PostCreateNewWallet(dataWallet, dataWalletAccount, dataTrans)

	if err != nil {
		res.RespCode = "9999"
		res.RespDesc = "error"
		res.RespDisplay = err.Error()
	} else {
		if result > 0 {

			res.RespCode = "1000"
			res.RespDesc = "Success"
			res.RespDisplay = "Success"
		} else {
			res.RespCode = "0001"
			res.RespDesc = "Failed"
			res.RespDisplay = ""
		}
	}

	return &res
}
