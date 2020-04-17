package internal

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

func NewDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("打开mysql链接失败: %w", err)
	}
	return db, nil
}

type DBData struct {
	DB     *sqlx.DB
	DBName string
}

var DBMap map[string]*DBData

var Src string
var Dest string

func Init() {
	DBMap = make(map[string]*DBData)
	Src = viper.GetString("src")
	Dest = viper.GetString("dest")

	srcDB, err := NewDB(viper.GetString(fmt.Sprintf("mysql.%s.dsn", Src)))
	if err != nil {
		panic(err)
	}
	DBMap[Src] = &DBData{
		DB:     srcDB,
		DBName: viper.GetString(fmt.Sprintf("mysql.%s.db", Src)),
	}
	destDB, err := NewDB(viper.GetString(fmt.Sprintf("mysql.%s.dsn", Dest)))
	if err != nil {
		panic(err)
	}
	DBMap[Dest] = &DBData{
		DB:     destDB,
		DBName: viper.GetString(fmt.Sprintf("mysql.%s.db", Dest)),
	}
}
