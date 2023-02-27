package service

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/maxxjula/wallet-service/repository"
)

func (s *Service) ApiPostDeductBalance(c echo.Context, req RequestDeductBalance) *ResponseDeductBalance {

	var res ResponseDeductBalance

	timenow := time.Now()
	data := repository.TbTransaction{
		Type:      "W",
		AccountNo: req.AccountNo,
		Amount:    req.Amount,
		TransDt:   timenow,
		CreatedDt: timenow,
		CreatedBy: "admin"}
	result, err := s.repository.PostDeductBalance(data)
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
