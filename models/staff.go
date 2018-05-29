package models

import (

	"fmt"
	"log"
	"database/sql"
)

type Staff struct {
	ID int64 `form:"-"`
	FirstName string `form: "first_name"`
	LastName string `form: "last_name"`
	Position string `form: "position"`
	Admin bool `form: "admin""`
}

var (
	staff [] Staff
	id int64
	firstName string
	lastName string
	position string
	admin bool
)


func ListStaff (id int64) (staff Staff) {
	db := Connect()

	err := db.QueryRow("SELECT * FROM staff where id=?", id).Scan(&staff.ID, &staff.FirstName, &staff.LastName, &staff.Position, &staff.Admin)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		log.Fatal("Unable to scan vars in List staff %s\n", err)
	}

	return staff
}

func ListAllStaff (staff [] Staff) {

}

func CreateStaff (staff Staff) (zoostaff Staff) {
	db := Connect()

	stmtIn, err := db.Prepare("INSERT into staff (first_name, last_name, position, admin) on VALUES (?,?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	defer stmtIn.Close()

	res, err := stmtIn.Exec(staff.FirstName, staff.LastName, staff.Position, staff.Admin);

	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	lastId, err := res.LastInsertId()

	staff = ListStaff(lastId)

	fmt.Println(affect)
	return zoostaff
}
