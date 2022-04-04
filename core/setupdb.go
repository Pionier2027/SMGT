/* models/のデータをセットアップする */

package core

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// ref: https://github.com/flipped-aurora/gin-vue-admin/blob/28eea6138a7ecae9d5eddca1dc5ca4e3b5c1007e/server/model/system/request/sys_init.go#L9
type InitDB struct {
	Host string
	User string
	Pass string
	Name string
}

// (テーブルではなく)データベースを作るためのdsnを取得する。
// ref : https://github.com/flipped-aurora/gin-vue-admin/blob/28eea6138a7ecae9d5eddca1dc5ca4e3b5c1007e/server/model/system/request/sys_init.go#L32
func (initDB *InitDB) EmptyDsn() string {
	url := fmt.Sprintf("host=%s user=%s password=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo", initDB.Host, initDB.User, initDB.Pass)
	return url
}

func (initDB *InitDB) CreateDatabase() error {
	dsn := initDB.EmptyDsn()
	createSql := "CREATE DATABASE " + initDB.Name
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func SetupDB() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	Host := os.Getenv("DB_HOST")
	User := os.Getenv("DB_USER")
	Pass := os.Getenv("DB_PASS")
	Name := os.Getenv("DB_NAME")
	initDB := InitDB{Host: Host, User: User, Pass: Pass, Name: Name}
	err := initDB.CreateDatabase()
	if err != nil {
		return err
	}
	return err
	// Init Tables
	// ref: https://github.com/flipped-aurora/gin-vue-admin/blob/28eea6138a7ecae9d5eddca1dc5ca4e3b5c1007e/server/service/system/sys_initdb_pgsql.go#L53

}
