package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func conn() (*sql.DB, error) {
	dataSource := "root:root@tcp(localhost:3306)/movies_db?parseTime=true"

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("could not ping the database: ", err)
		return nil, err
	}

	log.Println("connected")
	return db, nil
}

type movie struct {
	created_at  time.Time
	updated_at  time.Time
	title       string
	rating      float32
	awards      uint32
	releaseDate time.Time
	length      uint
	genre_id    uint32
}

func main() {
	db, err := conn()
	if err != nil {
		log.Fatal("could not open the conection: ", err)
	}

	defer db.Close()

	myMovie := &movie{
		created_at:  time.Now(),
		updated_at:  time.Now(),
		title:       "Capitão América",
		rating:      9.8,
		awards:      10,
		releaseDate: time.Now(),
		length:      117,
		genre_id:    1,
	}

	stmt, err := db.Prepare(`
    INSERT INTO movies (created_at, updated_at, title, rating, awards, release_date, length, genre_id) 
    VALUES (?, ?, ?, ?, ?, ?, ?, ?) 
  `)
	if err != nil {
		log.Println("failed to prepare: ", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		&myMovie.created_at,
		&myMovie.updated_at,
		&myMovie.title,
		&myMovie.rating,
		&myMovie.awards,
		&myMovie.releaseDate,
		&myMovie.length,
		&myMovie.genre_id,
	)
	if err != nil {
		log.Println("failed to exec: ", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Println("failed to get the last id: ", err)
	}

	fmt.Println("Last ID: ", lastID)
}
