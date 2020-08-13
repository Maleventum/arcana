package database

import (
	"log"

	"github.com/Maleventum/arcana/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DBDriver = "sqlite3"

type Database struct {
	DB *sqlx.DB
}

func New() *Database {
	db, err := sqlx.Open(DBDriver, "hackabot.db")
	if err != nil {
		log.Printf("ERROR: connecting to the database: %v\n", err)
		return nil
	}

	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS event (id INTEGER PRIMARY KEY, owner INTEGER, name TEXT, description TEXT)")
	_, err = statement.Exec()
	statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, cookie TEXT)")
	_, err = statement.Exec()
	if err != nil {
		log.Printf("error creating database: %v\n", err)
	}

	return &Database{DB: db}
}

func (s *Database) GetEvent(owner int64) []model.Event {
	events := []model.Event{}
	statement := "SELECT * FROM event"
	if owner > 0 {
		statement += " WHERE owner = ?"
	}

	err := s.DB.Select(&events, statement, owner)
	if err != nil {
		log.Println(err)
	}

	return events
}

func (s *Database) GetEventByID(id int64) []model.Event {
	events := []model.Event{}
	statement := "SELECT * FROM event  WHERE id = ?"

	err := s.DB.Select(&events, statement, id)
	if err != nil {
		log.Println(err)
	}

	return events
}

func (s *Database) CreateEvent(event *model.Event) error {
	insert, err := s.DB.Prepare("INSERT INTO event (owner, name, description) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("Error during insert event: %v\n", err)
		return err
	}
	// TODO: Return event identifier
	_, err = insert.Exec(event.Owner, event.Name, event.Description)

	return err
}

func (s *Database) UpdateEvent(event model.Event) error {
	statement := "UPDATE event SET "
	if event.Name != "" {
		statement += "name=? WHERE id=?"
		update, _ := s.DB.Prepare(statement)
		_, err := update.Exec(event.Name, event.ID)
		return err
	}
	if event.Description != "" {
		statement += "description=? WHERE id=?"
		update, _ := s.DB.Prepare(statement)
		_, err := update.Exec(event.Description, event.ID)
		return err
	}
	return nil
}

func (s *Database) DeleteEvent(id int64) error {
	delete, err := s.DB.Prepare("DELETE FROM event WHERE id=?")
	if err != nil {
		return err
	}
	_, err = delete.Exec(id)
	return err
}
