package repository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *DB {
	return &DB{db: db}
}

func (db *DB) GetAllWallet() ([]TbWallet, error) {
	var Res []TbWallet
	result := db.db.Table("wallet").
		Select(`*`).
		Order(`id`).
		Scan(&Res)
	return Res, result.Error
}

func (db *DB) PostCreateNewWallet(req TbWallet, reqWAcc TbWalletAccount, reqTrans TbTransaction) (int64, error) {

	result := db.db.Create(req)
	if result.RowsAffected > 0 {
		resultWAcc := db.db.Create(reqWAcc)
		if resultWAcc.RowsAffected > 0 {
			resultTrans, err := db.PostAddBalance(reqTrans, "new")
			return resultTrans, err
		}
		return resultWAcc.RowsAffected, resultWAcc.Error
	}
	return result.RowsAffected, result.Error
}

func (db *DB) GetWalletDetail(cid, accNo string) ([]TmpWalletDetail, error) {
	var Res []TmpWalletDetail

	result := db.db.Table("wallet_account as a").
		Select(`a.citizen_id,
		a.account_no,
		a.balance,
		a.status,
		a.remark,
		b.title,
		b.firstname,
		b.lastname`).
		Joins(`inner join wallet as b on a.citizen_id=b.citizen_id`).
		Where(`a.citizen_id = ?
		and a.account_no =?`, cid, accNo).
		Scan(&Res)
	return Res, result.Error
}

func (db *DB) PostAddBalance(req TbTransaction, status string) (int64, error) {

	result := db.db.Create(req)

	if result.RowsAffected > 0 {
		if status == "old" {
			resultUpdate := db.db.Exec(`UPDATE wallet_account as a,
		(SELECT balance as amt_balance FROM wallet_account WHERE account_no='` + req.AccountNo + `' AND STATUS='A') as b
		SET a.balance=b.amt_balance + ` + fmt.Sprintf("%.2f", req.Amount) + `,
		a.updated_dt='` + time.Now().Format("2006-01-02 15:04:05") + `',
		a.updated_by='admin'
		WHERE a.account_no='` + req.AccountNo + `' 
		AND STATUS='A'`)

			return resultUpdate.RowsAffected, resultUpdate.Error
		} else if status == "new" {
			resultUpdate := db.db.Exec(`UPDATE wallet_account as a
		SET a.balance=` + fmt.Sprintf("%.2f", req.Amount) + `,
		a.updated_dt='` + time.Now().Format("2006-01-02 15:04:05") + `',
		a.updated_by='admin'
		WHERE a.account_no='` + req.AccountNo + `' 
		AND STATUS='A'`)

			return resultUpdate.RowsAffected, resultUpdate.Error
		}

	}

	return result.RowsAffected, result.Error
}

func (db *DB) PostDeductBalance(req TbTransaction) (int64, error) {

	result := db.db.Create(req)

	if result.RowsAffected > 0 {
		resultUpdate := db.db.Exec(`UPDATE wallet_account as a,
		(SELECT balance as amt_balance FROM wallet_account WHERE account_no='` + req.AccountNo + `' AND STATUS='A') as b
		SET a.balance=b.amt_balance - ` + fmt.Sprintf("%.2f", req.Amount) + `,
		a.updated_dt='` + time.Now().Format("2006-01-02 15:04:05") + `',
		a.updated_by='admin'
		WHERE a.account_no='` + req.AccountNo + `' 
		AND STATUS='A'`)

		return resultUpdate.RowsAffected, resultUpdate.Error
	}

	return result.RowsAffected, result.Error
}

func (db *DB) PostWalletStatus(req TbWalletAccount) (int64, error) {

	result := db.db.Model(&req).Where(`account_no = ?`, req.AccountNo).Updates(map[string]interface{}{
		"status":     req.Status,
		"updated_dt": time.Now(),
		"updated_by": "admin"})

	return result.RowsAffected, result.Error
}
