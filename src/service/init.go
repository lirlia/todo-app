package service

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger "gorm.io/gorm/logger"
	"todo.app/conf"
	"todo.app/model"
)

var db *gorm.DB

func init() {

	dsn := conf.DB_USER + ":" + conf.DB_PASS + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	err := errors.New("")

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// トランザクションの中で整合性を担保するための設定
		// 今回は特定のレコードに対して複数の操作を行うことはないため性能向上を鑑みてtrueとする
		// https://gorm.io/docs/performance.html#Disable-Default-Transaction
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	db.Logger = db.Logger.LogMode(logger.Silent)

	sqlDB, err := db.DB()
	// SetMaxIdleConnsはアイドル状態のコネクションプール内の最大数を設定します
	sqlDB.SetMaxIdleConns(conf.DB_MAX_IDLE_CONN)
	// SetMaxOpenConnsは接続済みのデータベースコネクションの最大数を設定します
	sqlDB.SetMaxOpenConns(conf.DB_MAX_OPEN_CONN)
	// SetConnMaxLifetimeは再利用され得る最長時間を設定します
	sqlDB.SetConnMaxLifetime(conf.DB_MAX_LIFETIME)
	if err != nil {
		panic(err)
	}

	// テーブルの作成を行う
	type TableList = []model.Tables
	var t TableList
	t = append(t, &model.User{})
	t = append(t, &model.Task{})

	for _, v := range t {
		// テーブルの存在をチェックしない場合のみ作る
		if !db.Migrator().HasTable(v) {
			// テーブルの作成
			db.Migrator().CreateTable(v)
		}
	}

	fmt.Println("init data base ok")
}
