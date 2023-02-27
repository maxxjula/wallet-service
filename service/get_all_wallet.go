package service

import "github.com/labstack/echo/v4"

func (s *Service) ApiGetAllWallet(c echo.Context) *ResponseAllWallet {

	var res ResponseAllWallet
	var resValue []ValueAllWallet

	list, err := s.repository.GetAllWallet()

	if err != nil {
		res.RespCode = "9999"
		res.RespDesc = "error"
		res.RespDisplay = err.Error()
	} else {
		if len(list) > 0 {
			for _, v := range list {

				resValue = append(resValue, ValueAllWallet{
					CitizenId:   v.CitizenId,
					Title:       v.Title,
					Firstname:   v.Firstname,
					Lastname:    v.Lastname,
					Address:     v.Address,
					Subdistrict: v.Subdistrict,
					District:    v.District,
					Province:    v.Province,
					ZipCode:     v.ZipCode,
					Tel:         v.Tel,
				})
			}
			res.RespCode = "1000"
			res.RespDesc = "Success"
			res.RespDisplay = "Success"
			res.Values = resValue
		} else {
			res.RespCode = "0001"
			res.RespDesc = "Failed"
			res.RespDisplay = "find not found"
		}
	}

	return &res
}
