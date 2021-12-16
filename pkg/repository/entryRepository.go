package repository

import (
	"database/sql"
	"fmt"
	"github.com/abhishekambastha/personal-diary/pkg/domain"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"time"
)

func AddEntry(entry domain.Entry) domain.Entry {
	db, err := sql.Open("mysql", "root:diary@tcp(127.0.0.1:3306)/diary?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	result, derr := db.Exec(`INSERT INTO entries (created_timestamp, text, picture, tid) VALUES ( ?, ?, ?, ?)`,
		time.Now(), entry.Text, entry.Picture, entry.Tid)
	if derr != nil {
		log.Fatal(derr)
	}

	tempId, _ := result.LastInsertId()

	entry.Id = strconv.Itoa(int(tempId))
	return entry
}

func GetAllEntries() []domain.Entry {
	db, err := sql.Open("mysql", "root:diary@tcp(127.0.0.1:3306)/diary?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := `SELECT id, created_timestamp, text, picture, tid FROM entries`
	rows, dbErr := db.Query(query)
	defer rows.Close()

	if dbErr != nil {
		fmt.Println(dbErr)
	}

	var entries []domain.Entry
	for rows.Next() {
		var e domain.Entry
		err = rows.Scan(&e.Id, &e.Date, &e.Text, &e.Picture, &e.Tid)
		entries = append(entries, e)
	}

	return entries
}

func GetEntryById(id string) domain.Entry {
	db, err := sql.Open("mysql", "root:diary@tcp(127.0.0.1:3306)/diary?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := `SELECT id, created_timestamp, text, picture, tid FROM entries Where id = ?`
	idInt, _ := strconv.Atoi(id)

	var entry domain.Entry
	dbErr := db.QueryRow(query, idInt).Scan(&entry.Id, &entry.Date, &entry.Text, &entry.Picture, &entry.Tid)

	if dbErr != nil {
		fmt.Println(dbErr)
	}

	return entry
}
