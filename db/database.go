package db

import (
	"fmt"
	"log"
	"os"
	"rem/config"
	"rem/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 横着したのでそのうち直せ
var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dbConf, err := config.LoadDBConfig()
		if err != nil {
			log.Fatal("DB設定の読み込みに失敗:", err)
		}

		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Tokyo",
			dbConf.DB_HOST, dbConf.DB_USER, dbConf.DB_PASSWORD, dbConf.DB_NAME, dbConf.DB_PORT, dbConf.DB_SSLMODE,
		)
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("PostgreSQL接続に失敗:", err)
	}

	// テーブル作成
	err = DB.AutoMigrate(&models.Reminder{})
	if err != nil {
		log.Fatal("マイグレーションに失敗:", err)
	}
}
