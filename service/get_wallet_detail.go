package service

import (
	"github.com/labstack/echo/v4"
)

func (s *Service) ApiGetWalletDetail(c echo.Context, req RequestWalletDetail) *ResponseWalletDetail {
	var res ResponseWalletDetail
	var resValue []ValueWalletDetail

	list, err := s.repository.GetWalletDetail(req.CitizenId, req.AccountNo)

	if err != nil {
		res.RespCode = "9999"
		res.RespDesc = "error"
		res.RespDisplay = err.Error()

	} else {
		if len(list) > 0 {
			// if list != (repository.TmpWalletDetail{}) {
			// resValue = ValueWalletDetail{
			// 	AccountNo: list.AccountNo,
			// 	CitizenId: list.CitizenId,
			// 	Title:     list.Title,
			// 	Firstname: list.Firstname,
			// 	Lastname:  list.Lastname,
			// 	Balance:   list.Balance,
			// 	Status:    list.Status,
			// 	Remark:    list.Remark,
			// }

			for _, v := range list {

				resValue = append(resValue, ValueWalletDetail{
					AccountNo: v.AccountNo,
					CitizenId: v.CitizenId,
					Title:     v.Title,
					Firstname: v.Firstname,
					Lastname:  v.Lastname,
					Balance:   v.Balance,
					Status:    v.Status,
					Remark:    v.Remark,
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
