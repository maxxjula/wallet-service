# wallet-service
1. run file db.sql to create mySQL database 
2. run program with go run main.go
3. run API with postman 
3.1 API list all wallet
    http://localhost:8080/wallet-all method GET


3.2 API crate a new wallet
    http://localhost:8080/wallet-new method POST
    ex. request
    {
    "citizen_id" : "2209141648200",
    "title":"นางสาว",
    "firstname":"สมหญิง",
    "lastname": "ใจจันทร์",
    "address":"ถนนสุขุมวิท",
    "subdistrict":"พระโขนง",
    "district":"บางนา",
    "province":"กรุงเทพ",
    "tel":"104444",
    "zipcode": "54120",
    "account_no":"1111112200",
    "amount":800    
}

3.3 API get a wallet's detail
    http://localhost:8080/wallet-detail method GET
    ex. request
    {
    "account_no" : "1234567890",
    "citizen_id" : "2209141648120"
}

3.4 API add balance to a wallet
    http://localhost:8080/add-balance method  POST
    ex. request
{
    "account_no" : "4444444888",
    "amount":100   
}

3.5 API deduct balance from wallet
    http://localhost:8080/deduct-balance method POST
    ex. request
    {
    "account_no" : "4444444888",
    "amount":10.32
}

3.6 API deactivate & activate a wallet
    http://localhost:8080/wallet-status method POST
    {
    "account_no" : "1234444444",
    "status": "A"
}