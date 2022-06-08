package database

import (
	"fmt"
	"github.com/chinaboard/habada/storage/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

type MysqlStorage struct {
	db *gorm.DB
}

func New() *MysqlStorage {
	user, password, host, database := os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		database)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	s, _ := db.DB()
	s.SetMaxIdleConns(10)
	s.SetMaxOpenConns(100)
	s.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.HabadaItem{})

	return &MysqlStorage{db: db}
}

func (m *MysqlStorage) Get(tinyUrl string) (longUrl string, err error) {
	var item model.HabadaItem
	result := m.db.Where("tiny_url = ?", tinyUrl).First(&item)
	if result.Error != nil {
		return "", result.Error
	}
	return item.LongUrl, nil
}

func (m *MysqlStorage) Set(tinyUrl, longUrl string) (success bool, err error) {
	result := m.db.Create(&model.HabadaItem{TinyUrl: tinyUrl, LongUrl: longUrl})
	return result.Error == nil, err
}
