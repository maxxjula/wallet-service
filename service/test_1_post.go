package service

import "github.com/labstack/echo/v4"

func (s *Service) ApiTest1(c echo.Context, req ReqTest1) *ResTest1 {
	var res ResTest1
	res.RespCode = "1000"
	res.RespDesc = "Success"
	res.RespDisplay = "Success"
	return &res
}
