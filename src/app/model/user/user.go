package user

import (
	"core/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id   int
	Name string
}

func GetUserById(id int) (*User, error) {
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

	u := User{}
	err = db.QueryRow("select * from user where id = ?", id).Scan(&u.Id, &u.Name)
	if err != nil {
		return nil, err
	}
	fmt.Println(u)

	return &u, nil
}
