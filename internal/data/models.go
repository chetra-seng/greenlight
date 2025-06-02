package data

import (
	"database/sql"
	"errors"
)

type Models struct {
	Movies MovieModel
}

var (
	ErrRecordNotFound = errors.New("record not found")
)

func NewModel(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
