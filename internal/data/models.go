package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Watches WatchModel
	Users   UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Watches: WatchModel{DB: db},
		Users:   UserModel{DB: db},
	}
}

//func NewMockModels() Models {
//	return Models{
//		Watches: MockWatchModel{},
//	}
//}
