package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"u_demo/conf"
)

// gorm 实现版本
//var DB *gorm.DB
//
//func Init_Mysql(cfg *conf.MySQLConfig) error {
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Printf("gorm mysql not open, err is %v", err)
//		return err
//	}
//
//	sqlDB, err := db.DB() // database/sql.DB
//	if err != nil {
//		fmt.Printf(" 额外的连接配置失败， err is %v\n", err)
//		return err
//	}
//	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
//	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
//	sqlDB.SetConnMaxLifetime(time.Hour)
//	fmt.Println("mysql 配置成功")
//	DB = db
//	return nil
//}
//
//// 关闭数据库连接
//func CloseDatabase() {
//	sqlDB, err := DB.DB() // 获取原始 *sql.DB 对象
//	if err != nil {
//		zap.L().Error("failed to get SQL DB from GORM")
//	}
//	err = sqlDB.Close() // 关闭数据库连接
//	if err != nil {
//		zap.L().Error("failed to close database connection")
//	}
//}

var db *sqlx.DB

// Init 初始化MySQL连接
func Init(cfg *conf.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
