package models_x

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gowebsocket/utils_x"
	"log"
	"time"
)

var db = InitGorm()

/**
* @Author junfenghe
* @Description 初始化 Gorm 连接池
* @Date 2021-03-20 8:24
* @Param
* @return
**/
func InitGorm() *gorm.DB {
	user := "cbbpa"
	password := "cbbpa@456!"
	host := "cdb-7b2psey2.bj.tencentcdb.com"
	port := "10178"
	database := "cbbpa_im"
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=true&loc=Local"
	log.Println(dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:        "mysql",
		DSN:               dsn,
		DefaultStringSize: 255,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	sqlDB.SetMaxIdleConns(12)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute * 1)
	return db
}

/**
* @Author junfenghe
* @Description Func of BeforeCrete
* @Date 2021-03-20 8:17
* @Param
* @return
**/
func BeforeCreateUpdateID(tx *gorm.DB) error {
	// 创建之前生成 uuid
	uuid, err := utils_x.GenerateUUID()
	if err != nil {
		log.Println(err)
		return err
	}
	tx.Statement.SetColumn("ID", uuid.(string))
	return nil
}
