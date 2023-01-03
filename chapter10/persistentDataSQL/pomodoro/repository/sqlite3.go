package repository

import (
	"database/sql"
	"sync"
	"time"

	"github.com/eduardohitek/powerful-cli/chapter9/pomo/pomodoro"
	_ "github.com/mattn/go-sqlite3"
)

const (
	createTableInterval string = `CREATE TABLE IF NOT EXISTS "interval" (
		"id"    INTEGER,
		"start_time"    DATETIME NOT NULL,
		"planned_duration"      INTEGER DEFAULT 0,
		"actual_duration"       INTEGER DEFAULT 0,
		"category"  TEXT NOT NULL,
		"state" INTEGER DEFAULT 1,
		PRIMARY KEY("id")
	);`
)

type dbRepo struct {
	db *sql.DB
	sync.RWMutex
}

func NewSQLite3Repo(dbfile string) (*dbRepo, error) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetMaxOpenConns(1)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createTableInterval)
	if err != nil {
		return nil, err
	}
	return &dbRepo{
		db: db,
	}, nil

}

func (r *dbRepo) Create(i pomodoro.Interval) (int64, error) {
	r.Lock()
	defer r.Unlock()

	insSmt, err := r.db.Prepare("INSERT INTO interval VALUES(NULL, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer insSmt.Close()

	res, err := insSmt.Exec(i.StartTime, i.PlannedDurarion, i.ActualDuration, i.Category, i.State)
	if err != nil {
		return 0, err
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
