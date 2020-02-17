package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB database

type database interface {
	GetUsers(ctx context.Context) (users []User, err error)
	NewUser(input User) (user User, err error)
	NewHome(input Home) (home Home, err error)
}

// DBInfo groups DB info
type DBInfo struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

const UserSchema = `
CREATE TABLE IF NOT EXISTS user (
	id int(11) NOT NULL auto_increment,
	name  varchar(100)  NULL default '',
	password  varchar(100)  NULL default '',
	email  varchar(100)  NULL default '',
	address  varchar(100)  NULL default '',
	age int(11)  NULL default 0,    
  
	PRIMARY KEY  (id)
);
`
const HomeSchema = `
CREATE TABLE IF NOT EXISTS home (
	id int(11) NOT NULL auto_increment,
	serial  varchar(100)  NULL default '',
	user_id int,

	PRIMARY KEY  (id),
	FOREIGN KEY (user_id) REFERENCES user(id)
  );
`

const HomeSensorsSchema = `
CREATE TABLE IF NOT EXISTS home_sensors (
	id int,
	name  varchar(100)  NULL default '',
	home_id int,

	PRIMARY KEY  (id),
	FOREIGN KEY (home_id) REFERENCES home(id)
  );
`

func ConfigureSQL(dbinfo DBInfo) (err error) {
	db, err := ConnectDb(dbinfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	DB = NewSQLDatabase(db)
	db.MustExec(UserSchema)
	db.MustExec(HomeSchema)
	db.MustExec(HomeSensorsSchema)

	log.Println("Successfully connected to SQL database")
	return
}

// ConnectDb connects to the DB
func ConnectDb(dbinfo DBInfo) (*sqlx.DB, error) {
	fmt.Printf("Database connected %s at %s:%d\n", dbinfo.Name, dbinfo.Host, dbinfo.Port)

	dbsql := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbinfo.User, dbinfo.Pass, dbinfo.Host, dbinfo.Port, dbinfo.Name)
	return sqlx.Open("mysql", dbsql)
}
