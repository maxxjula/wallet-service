package service

import (
	"github.com/labstack/echo/v4"
	"github.com/maxxjula/wallet-service/repository"
)

func (s *Service) ApiPostWalletStatus(c echo.Context, req RequestWalletStatus) *ResponseWalletStatus {

	var res ResponseWalletStatus
	var data repository.TbWalletAccount

	if req.Status == "A" || req.Status == "Activate" || req.Status == "activate" {
		data = repository.TbWalletAccount{
			AccountNo: req.AccountNo,
			Status:    "A"}

	} else if req.Status == "D" || req.Status == "Deactivate" || req.Status == "deactivate" {
		data = repository.TbWalletAccount{
			AccountNo: req.AccountNo,
			Status:    "D"}
	} else {
		res.RespCode = "0001"
		res.RespDesc = "Failed"
		res.RespDisplay = "status not match"
		return &res
	}
	
	result, err := s.repository.PostWalletStatus(data)
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
