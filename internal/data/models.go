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
	Watches interface {
		Insert(watch *Watch) error
		Get(id int64) (*Watch, error)
		Update(watch *Watch) error
		Delete(id int64) error
		GetAll(brand string,
			dialColor string,
			strapType string,
			diameter int8,
			energy string,
			gender string,
			priceRange []int,
			filters Filters) ([]*Watch, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Watches: WatchModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Watches: MockWatchModel{},
	}
}
