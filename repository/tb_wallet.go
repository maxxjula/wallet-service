package repository

import "time"

type TbWallet struct {
	CitizenId   string    `gorm:"column:citizen_id" json:"citizen_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Firstname   string    `gorm:"column:firstname" json:"firstname"`
	Lastname    string    `gorm:"column:lastname" json:"lastname"`
	Address     string    `gorm:"column:address" json:"address"`
	Subdistrict string    `gorm:"column:subdistrict" json:"subdistrict"`
	District    string    `gorm:"column:district" json:"district"`
	Province    string    `gorm:"column:province" json:"province"`
	ZipCode     string    `gorm:"column:zipcode" json:"zipcode"`
	Tel         string    `gorm:"column:tel" json:"tel"`
	CreatedDt   time.Time `gorm:"column:created_dt" json:"created_dt"`
	CreatedBy   string    `gorm:"column:created_by" json:"created_by"`
	UpdatedDt   time.Time `gorm:"column:updated_dt" json:"updated_dt"`
	UpdatedBy   string    `gorm:"column:updated_by" json:"updated_by"`
}

func (TbWallet) TableName() string {
	return "wallet"
}
