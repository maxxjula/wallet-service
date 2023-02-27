package service

import "time"

type Register struct {
	UserName string    `่json:"username"`
	Password string    `่json:"password"`
	Name     string    `่json:"name"`
	Age      int64     `่json:"age"`
	Brithday time.Time `่json:"brithday"`
}

type Response struct {
	Name string `่json:"name,omitempty"`
	// Age  *int   `่json:"age,omitempty"`
}
type Request struct {
	UserName string `่json:"username"`
	Password string `่json:"password"`
}

//++++++++++++++++++++++++++++++++

type ResponseAllWallet struct {
	RespCode    string           `json:"respCode"`
	RespDesc    string           `json:"respDesc"`
	RespDisplay string           `json:"respDisplay"`
	Values      []ValueAllWallet `json:"values"`
}

type ValueAllWallet struct {
	CitizenId   string `json:"citizen_id"`
	Title       string `json:"title"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Address     string `json:"address"`
	Subdistrict string `json:"subdistrict"`
	District    string `json:"district"`
	Province    string `json:"province"`
	ZipCode     string `json:"zipcode"`
	Tel         string `json:"tel"`
}

type RequestNewWallet struct {
	CitizenId   string  `validate:"nonzero,max=13,regexp=^[0-9]*$" json:"citizen_id"`
	Title       string  `validate:"nonzero,max=20,regexp=^[\u0E00-\u0E7Fa-zA-Z0-9-_:\\,.#$+/\\(\\)\\/\\s]*$" json:"title"`
	Firstname   string  `validate:"nonzero,max=50,regexp=^[\u0E00-\u0E7Fa-zA-Z0-9-_:\\,.#$+/\\(\\)\\/\\s]*$" json:"firstname"`
	Lastname    string  `validate:"nonzero,max=50,regexp=^[\u0E00-\u0E7Fa-zA-Z0-9-_:\\,.#$+/\\(\\)\\/\\s]*$" json:"lastname"`
	Address     string  `validate:"nonzero,max=200,regexp=^[\u0E00-\u0E7Fa-zA-Z0-9-_:\\,.#$+/\\(\\)\\/\\s]*$" json:"address"`
	Subdistrict string  `validate:"nonzero,max=100,regexp=^[\u0E00-\u0E7Fa-zA-Z0-9-_:\\,.#$+/\\(\\)\\/\\s]*$" json:"subdistrict"`
	District    string  `validate:"nonzero,max=100,regexp=^[\u0E00-\u0E7Fa-zA-Z0-9-_:\\,.#$+/\\(\\)\\/\\s]*$" json:"district"`
	Province    string  `validate:"nonzero,max=100,regexp=^[\u0E00-\u0E7Fa-zA-Z0-9-_:\\,.#$+/\\(\\)\\/\\s]*$" json:"province"`
	ZipCode     string  `validate:"nonzero,max=5,regexp=^[0-9]*$" json:"zipcode"`
	Tel         string  `validate:"nonzero,max=20,regexp=^[0-9]*$" json:"tel"`
	AccountNo   string  `validate:"nonzero,max=10,regexp=^[0-9]*$" json:"account_no"`
	Remark      string  `validate:"max=250,regexp=^[\u0E00-\u0E7Fa-zA-Z0-9-_:\\,.#$+/\\(\\)\\/\\s]*$" json:"remark"`
	Amount      float64 `validate:"nonzero,max=99999999.9999" json:"amount"`
}

type ResponseNewWallet struct {
	RespCode    string `json:"respCode"`
	RespDesc    string `json:"respDesc"`
	RespDisplay string `json:"respDisplay"`
}

type RequestWalletDetail struct {
	CitizenId string `validate:"nonzero,max=13,regexp=^[0-9]*$" json:"citizen_id"`
	AccountNo string `validate:"nonzero,max=10,regexp=^[0-9]*$" json:"account_no"`
}

type ResponseWalletDetail struct {
	RespCode    string              `json:"respCode"`
	RespDesc    string              `json:"respDesc"`
	RespDisplay string              `json:"respDisplay"`
	Values      []ValueWalletDetail `json:"values"`
}

type ValueWalletDetail struct {
	CitizenId string `json:"citizen_id"`
	AccountNo string `json:"account_no"`
	Title     string `json:"title"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Balance   string `json:"balance"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
}

type RequestAddBalance struct {
	AccountNo string  `validate:"nonzero,max=10,regexp=^[0-9]*$" json:"account_no"`
	Amount    float64 `validate:"nonzero,max=99999999.9999" json:"amount"`
}

type ResponseAddBalance struct {
	RespCode    string `json:"respCode"`
	RespDesc    string `json:"respDesc"`
	RespDisplay string `json:"respDisplay"`
}

type RequestDeductBalance struct {
	AccountNo string  `validate:"nonzero,max=10,regexp=^[0-9]*$" json:"account_no"`
	Amount    float64 `validate:"nonzero,max=99999999.9999" json:"amount"`
}

type ResponseDeductBalance struct {
	RespCode    string `json:"respCode"`
	RespDesc    string `json:"respDesc"`
	RespDisplay string `json:"respDisplay"`
}

type RequestWalletStatus struct {
	AccountNo string `validate:"nonzero,max=10,regexp=^[0-9]*$" json:"account_no"`
	Status    string `validate:"nonzero,max=10,regexp=^[\u0E00-\u0E7Fa-zA-Z0-9-_:\\,.#$+/\\(\\)\\/\\s]*$" json:"status"`
}

type ResponseWalletStatus struct {
	RespCode    string `json:"respCode"`
	RespDesc    string `json:"respDesc"`
	RespDisplay string `json:"respDisplay"`
}
