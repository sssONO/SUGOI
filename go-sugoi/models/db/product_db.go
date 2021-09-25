package db

import (
	// フォーマットI/O
	"fmt"
	// Go言語のORM
	"github.com/jinzhu/gorm"
	// エンティティ(データベースのテーブルの行に対応)
	entity "main/models/entity"
)

// DB接続する
func open() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "1225"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "sugoi"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(DBMS, CONNECT)
	// 詳細なログを表示
	db.LogMode(true)
	fmt.Println("db connected: ", &db)

	if err != nil {
		panic(err.Error())
	}
	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	// 詳細なログを表示
	db.LogMode(true)
	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)
	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&entity.Product{})
	fmt.Println("db connected: ", &db)
	return db
}

// InsertProduct は 商品テーブルにレコードを追加する
func InsertProduct(registerProduct *entity.Product) {
	db := open()
	db.Create(&registerProduct)
	defer db.Close()
}
