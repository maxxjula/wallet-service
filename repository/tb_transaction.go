package repository

import "time"

type TbTransaction struct {
	TransNo   string    `gorm:"column:trans_no" json:"trans_no"`
	Type      string    `gorm:"column:type" json:"type"`
	AccountNo string    `gorm:"column:account_no" json:"account_no"`
	Amount    float64    `gorm:"column:amount" json:"amount"`
	TransDt   time.Time `gorm:"column:trans_dt" json:"trans_dt"`
	CreatedDt time.Time `gorm:"column:created_dt" json:"created_dt"`
	CreatedBy string    `gorm:"column:created_by" json:"created_by"`
	UpdatedDt time.Time `gorm:"column:updated_dt" json:"updated_dt"`
	UpdatedBy string    `gorm:"column:updated_by" json:"updated_by"`
}

func (TbTransaction) TableName() string {
	return "transaction"
}
