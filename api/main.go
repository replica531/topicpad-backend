package main

import (
  "fmt"
  "time"

  "github.com/gin-gonic/gin"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  sqlConnect()

  router := gin.Default()
  router.LoadHTMLGlob("templates/*.html")

  router.GET("/", func(ctx *gin.Context){
    ctx.HTML(200, "index.html", gin.H{})
  })

  router.Run()
}

func sqlConnect() (database *gorm.DB) {
  USER := "go_test"
  PASSWORD := "password"
  PROTOCOL := "tcp(mysql:3306)"
  DBNAME := "go_database"

  dsn := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	fmt.Println(dsn)
  count := 0
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

  return db
}
