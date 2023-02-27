package repository

type TmpWalletDetail struct {
	AccountNo string `gorm:"column:account_no" json:"account_no"`
	CitizenId string `gorm:"column:citizen_id" json:"citizen_id"`
	Title     string `gorm:"column:title" json:"title"`
	Firstname string `gorm:"column:firstname" json:"firstname"`
	Lastname  string `gorm:"column:lastname" json:"lastname"`
	Balance   string `gorm:"column:balance" json:"balance"`
	Status    string `gorm:"column:status" json:"status"`
	Remark    string `gorm:"column:remark" json:"remark"`
}
