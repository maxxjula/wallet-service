package service

import (
	"github.com/labstack/echo/v4"
	"github.com/maxxjula/wallet-service/repository"
)

type Servicer interface {
	ApiGetAllWallet(c echo.Context) *ResponseAllWallet
	ApiPostCreatNewWallet(c echo.Context, req RequestNewWallet) *ResponseNewWallet
	ApiGetWalletDetail(c echo.Context, req RequestWalletDetail) *ResponseWalletDetail
	ApiPostAddBalance(c echo.Context, req RequestAddBalance) *ResponseAddBalance
	ApiPostDeductBalance(c echo.Context, req RequestDeductBalance) *ResponseDeductBalance
	ApiPostWalletStatus(c echo.Context, req RequestWalletStatus) *ResponseWalletStatus
}

type Dber interface {
	GetAllWallet() ([]repository.TbWallet, error)
	PostCreateNewWallet(req repository.TbWallet, reqWAcc repository.TbWalletAccount, reqTrans repository.TbTransaction) (int64, error)
	GetWalletDetail(cid, accNo string) ([]repository.TmpWalletDetail, error)
	PostAddBalance(req repository.TbTransaction, status string) (int64, error)
	PostDeductBalance(req repository.TbTransaction) (int64, error)
	PostWalletStatus(req repository.TbWalletAccount) (int64, error)
}

type Service struct {
	repository Dber
}

func NewService(db Dber) *Service {
	return &Service{repository: db}
}
