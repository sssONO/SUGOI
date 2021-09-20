package main

import (
	//無視してよい。
	"fmt"
	"time"

	"log"
	controller "main/controller"

	//"github.com/labstack/echo/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "1225"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "sugoi"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	// 詳細なログを表示
	db, err := gorm.Open(DBMS, CONNECT)

	db.LogMode(true)
	fmt.Println("db connected: ", &db)
	return db, err

}

func serve() {
	//ginの基礎
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://localhost:8080/yamabiko",
			"http://localhost:8000/yamabiko",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Content-Type",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// ルーターの設定
	// 商品情報をDBへ登録する
	engine.POST("/addProduct", controller.AddProduct)
	if err := engine.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}

func main() {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	serve()
}
