package main

import (
	"fmt"
	"golang-project/internal/model/task"
	"golang-project/internal/server"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	dmysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func exitIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	err := godotenv.Load()
	exitIfError(err)

	jst, err := time.LoadLocation(os.Getenv("TZ"))
	exitIfError(err)

	// TODO: user/passなどもENV化
	db, err := openGORM(os.Getenv("DB_NAME"), "root", "example", "0.0.0.0", os.Getenv("MYSQL_PORT_HOST"), jst)
	exitIfError(err)

	err = db.AutoMigrate(&task.Task{})
	exitIfError(err)

	e, err := server.NewServer(db)
	exitIfError(err)

	// TODO: PORTもENV化
	e.Logger.Fatal(e.Start(":9000"))
}

func openGORM(dbName, user, pass, host, port string, loc *time.Location) (*gorm.DB, error) {
	m := mysql.New(mysql.Config{
		DSNConfig: &dmysql.Config{
			User:      user,
			Passwd:    pass,
			DBName:    dbName,
			Addr:      fmt.Sprintf("%s:%s", host, port),
			Net:       "tcp",
			Loc:       loc,
			ParseTime: true,
		},
		DontSupportRenameIndex:        true,
		DontSupportRenameColumn:       true,
		DontSupportNullAsDefaultValue: true,
		DontSupportRenameColumnUnique: true,
	})

	db, err := gorm.Open(m, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
