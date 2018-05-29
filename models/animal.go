package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

type  ZooAnimal struct {
	ID int64 `form:"-"`
	AnimalName string `form:"animal_name"`
}

var (
	zooanimal [] ZooAnimal
	animalId int64
	animalName string
)

func Connect() *sql.DB {
	fmt.Println("In Connect")

	dbconfig := beego.AppConfig.String("db::user") + ":" + beego.AppConfig.String("db::pass") +
		"@tcp(" + beego.AppConfig.String("db::host") + ":" + beego.AppConfig.String("db::port") +
		")"

	db, err := sql.Open("mysql", dbconfig + "/" + beego.AppConfig.String("db::name"))
	if err != nil {
		log.Printf("Unable to prep db abstraction %s\n", err)

	}

	if err = db.Ping(); err != nil {
		log.Printf("Database Error %s\n", err)
	}

	return db
}

func ListAnimal (animalId int64) (zooanimal ZooAnimal) {

	db := Connect()

	err := db.QueryRow("SELECT * FROM animals where id=?", id).Scan(&zooanimal.ID, &zooanimal.AnimalName)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}

		log.Fatal("Unable to scan vars in List animal %s\n", err)
	}

	return animal

	return
}

func ListAnimals (animals []ZooAnimal) {

}

func CreateAnimal (animal ZooAnimal) (zooanimal ZooAnimal) {
	db := Connect()

	stmtIn, err := db.Prepare("INSERT into zoo_animals (animal_name)on) VALUES (?)")

	if err != nil {
		log.Fatal(err)
	}

	defer stmtIn.Close()

	res, err := stmtIn.Exec(animal.AnimalName);

	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	lastId, err := res.LastInsertId()

	zooanimal = ListAnimal(lastId)

	return zooanimal
}
