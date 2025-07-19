package data

import (
	"database/sql"
	"errors"
)

type Models struct {
	Movies MovieModel
	Users  UserModel
	Tokens TokenModel
}

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

func NewModel(db *sql.DB) Models {
	return Models{
		Users:  UserModel{DB: db},
		Movies: MovieModel{DB: db},
    Tokens: TokenModel{DB: db},
	}
}
