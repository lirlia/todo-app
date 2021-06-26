package model

// テーブル名を変換するために定義
type Tables interface {
	TableName() string
}
