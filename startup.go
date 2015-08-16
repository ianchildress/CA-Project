package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type server struct {
	Type string
	Ip   string
	Port string
}

func startup() {
	var db *sqlx.DB

	// This is a sloppy way to check if a database exists and create if it doesn't.
	//  Sloppy due to time restraints.
	db, err := sqlx.Open("mysql", "root:civis@tcp(civis-mysql:3306)/docker")
	err = db.Ping()
	if err != nil {
		initDB()
	}

	// This is a sloppy way to check if the tables were created properly, and create
	// them if they are not. Sloppy due to time restraints.
	_, err = db.Queryx("SELECT * FROM servers")
	if err != nil {
		initTables()
	}

	db.Close()
}

func initDB() {
	var db *sqlx.DB
	db, err := sqlx.Open("mysql", "root:civis@tcp(civis-mysql:3306)/")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE docker")
	if err != nil {
		panic(err)
	}
}

func initTables() {
	var db *sqlx.DB
	db, err := sqlx.Open("mysql", "root:civis@tcp(civis-mysql:3306)/docker")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var schema = `
		CREATE TABLE containers (
		    container_id varchar(32),
		    image varchar(32),
		    command varchar(32),
		    created timestamp,
		    status varchar(32),
		    ports varchar(32),
		    names varchar(32));`

	db.MustExec(schema)

	c := container{
		Container_id: "dev1",
		Image:        "mysql",
		Command:      "/entrypoint.sh",
		Created:      time.Now(),
		Status:       "Up 13 minutes",
		Ports:        "3306/tcp",
		Names:        "civis-mysql",
	}

	// Insert test data
	tx := db.MustBegin()
	r, err := tx.NamedExec("INSERT INTO containers (container_id, image, command, created, status, ports, names) VALUES (:container_id, :image, :command, :created, :status, :ports, :names)", &c)
	if err != nil {
		log.Panic(err)
	} else {
		log.Println(r)
	}
	tx.Commit()

	return
}
