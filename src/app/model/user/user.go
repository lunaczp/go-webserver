package user

import (
	"core/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id int
	Name string
}

func getUserById(id int) *User {
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		config.Config().DB.User,
		config.Config().DB.Pass,
		config.Config().DB.Url,
		config.Config().DB.Name,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//todo fetch
	return &User{}

}
