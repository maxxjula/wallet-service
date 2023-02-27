package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"github.com/maxxjula/wallet-service/repository"
	"github.com/maxxjula/wallet-service/service"
	"github.com/spf13/viper"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func init() {
	runtime.GOMAXPROCS(1)
	initViper()
}

func main() {
	var (
		db = newMariaConn()
	)
	r := echo.New()

	sers := service.NewHandle(service.NewService(repository.NewRepository(db)))
	// r.GET("/login", sers.CallLogin)
	r.GET("/wallet-all", sers.ExecuteGetAllWallet)
	r.POST("/wallet-new", sers.ExecuteCreateNewWallet)
	r.GET("/wallet-detail", sers.ExecuteGetWalletDetail)
	r.POST("/add-balance", sers.ExecuteAddBalance)
	r.POST("/deduct-balance", sers.ExecuteDeductBalance)
	r.POST("/wallet-status", sers.ExecuteWalletStatus)
	r.Logger.Fatal(r.Start(":"+viper.GetString("app.port")+""))

}
func initViper() {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config: %s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func newMariaConn() *gorm.DB {
	// dsn := "root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Printf("db error %s\n", err)
	// }
	// return db

	conf := mysql.Config{
		DBName:               viper.GetString("mysql.database"),
		User:                 viper.GetString("mysql.username"),
		Passwd:               viper.GetString("mysql.password"),
		Net:                  "tcp",
		Addr:                 viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port"),
		AllowNativePasswords: true,
		Timeout:              viper.GetDuration("mysql.timeout"),
		ReadTimeout:          viper.GetDuration("mysql.readtimeout"),
		WriteTimeout:         viper.GetDuration("mysql.writetimeout"),
		ParseTime:            viper.GetBool("mysql.parsetime"),
		Loc:                  time.Local,
	}

	fmt.Print("[CONFIG] repositories connection: ", strings.ReplaceAll(conf.FormatDSN(), conf.Passwd, "********"))

	conn, err := gorm.Open("mysql", conf.FormatDSN())
	if err != nil {

		fmt.Printf("cannot open mysql connection:%s", err)
	}

	conn.DB().SetMaxIdleConns(viper.GetInt("mysql.maxidle"))
	conn.DB().SetMaxOpenConns(viper.GetInt("mysql.maxopen"))
	conn.DB().SetConnMaxLifetime(viper.GetDuration("mysql.maxlifetime"))
	conn.LogMode(viper.GetBool("mysql.debug"))

	return conn
}
