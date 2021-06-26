package conf

import "time"

const (
	DB_NAME            = "todo"
	DB_USER            = "root"
	DB_PASS            = "password"
	DB_HOST            = "localhost"
	DB_PORT            = "3306"
	DB_USER_TABLE_NAME = "user"
	DB_TASK_TABLE_NAME = "task"

	DB_INSERT_BATCHSIZE = 3000

	// SetMaxIdleConnsはアイドル状態のコネクションプール内の最大数を設定します
	DB_MAX_IDLE_CONN = 10

	// SetMaxOpenConnsは接続済みのデータベースコネクションの最大数を設定します
	DB_MAX_OPEN_CONN = 100

	// SetConnMaxLifetimeは再利用され得る最長時間を設定します
	DB_MAX_LIFETIME = time.Hour
)
