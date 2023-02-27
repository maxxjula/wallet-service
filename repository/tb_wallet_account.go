package repository

import "time"

type TbWalletAccount struct {
	CitizenId string    `gorm:"column:citizen_id" json:"citizen_id"`
	AccountNo string    `gorm:"column:account_no" json:"account_no"`
	Balance   float64   `gorm:"column:balance" json:"balance"`
	Status    string    `gorm:"column:status" json:"status"`
	Remark    string    `gorm:"column:remark" json:"remark"`
	CreatedDt time.Time `gorm:"column:created_dt" json:"created_dt"`
	CreatedBy string    `gorm:"column:created_by" json:"created_by"`
	UpdatedDt time.Time `gorm:"column:updated_dt" json:"updated_dt"`
	UpdatedBy string    `gorm:"column:updated_by" json:"updated_by"`
}

func (TbWalletAccount) TableName() string {
	return "wallet_account"
}
