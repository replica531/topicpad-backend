package db

import (
	"time"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"topicpad-api/model"
)

var (
	db *gorm.DB
	err error
)

func Init() {
  USER := "topicpad_test"
  PASSWORD := "password"
  PROTOCOL := "tcp(mysql:3306)"
  DBNAME := "topicpad_database"

  dsn := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	fmt.Println(dsn)
  count := 0
  db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    for {
      if err == nil {
        fmt.Println("")
        break
      }
      time.Sleep(time.Second)
      count++
      if count > 180 {
        fmt.Println("")
        fmt.Println("DB接続失敗")
        panic(err)
      }
			fmt.Printf("retry(%d回目)\n", count)
      db, err =  gorm.Open(mysql.Open(dsn), &gorm.Config{})
    }
  }
  fmt.Println("DB接続成功")

	autoMigration()
	// db.Create(&model.User{Kind: 1})
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func autoMigration() {
	db.Migrator().CreateTable(&model.User{})
	// db.Migrator().CreateTable(&model.Topic{})
	db.Migrator().DropTable(&model.Topic{})
	fmt.Println(db.Migrator().CurrentDatabase())
}
