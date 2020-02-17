package main

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

// User defines the user model
type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Password string `json:"pass" db:"password"`
	Email    string `json:"email" db:"email"`
	Address  string `json:"address" db:"address"`
	Age      int    `json:"age" db:"age"`
	Homes    []Home `json:"homes"`
}

// Home defines the home model
type Home struct {
	ID      int      `json:"id" db:"id"`
	Serial  string   `json:"serial" db:"serial"`
	Sensors []Sensor `json:"sensors"`
	UserID  int      `json:"user_id" db:"user_id"`
}

// Sensor defines the sensor data
type Sensor struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	HomeID int    `json:"-" db:"home_id"`
}

type sqldatabase struct {
	db *sqlx.DB
}

// NewSQLDatabase returns a sqldatabase struct to implement database interface
func NewSQLDatabase(db *sqlx.DB) database {
	return sqldatabase{
		db: db,
	}
}

func (s sqldatabase) GetUsers(ctx context.Context) (users []User, err error) {
	s.db.SelectContext(ctx, &users, `select * from user;`)
	for i, u := range users {
		var homes []Home
		err = s.db.SelectContext(ctx, &homes, "SELECT * FROM home WHERE user_id = ?", u.ID)
		users[i].Homes = homes
	}
	return
}

func (s sqldatabase) NewUser(input User) (user User, err error) {
	// TODO: pass the sql result to user output function
	user = input
	userResult := s.db.MustExecContext(context.TODO(),
		"INSERT INTO user(name, password, email , address , age) VALUES( ?, ?, ?, ?, ? )",
		input.Name,
		input.Password,
		input.Email,
		input.Address,
		input.Age,
	)

	userRecordID, err := userResult.LastInsertId()
	if err != nil {
		return
	}

	user.ID = int(userRecordID)

	return
}

func (s sqldatabase) NewHome(input Home) (home Home, err error) {
	// TODO: pass the sql result to user output function
	home = input
	homeResult := s.db.MustExecContext(context.TODO(),
		"INSERT INTO home(serial, user_id) VALUES( ?, ? )",
		input.Serial,
		input.UserID,
	)

	homeRecordID, err := homeResult.LastInsertId()
	if err != nil {
		return
	}

	home.ID = int(homeRecordID)

	for _, i := range input.Sensors {
		log.Println(i)
		fn := s.db.MustExecContext(context.TODO(),
			"INSERT INTO home_sensors(id , name, home_id) VALUES( ?, ?, ? )",
			i.ID,
			i.Name,
			int(homeRecordID),
		)
		if _, fnerr := fn.LastInsertId(); fnerr != nil {
			err = fnerr
			return
		}
	}

	return
}
