package data

import (
	"database/sql"
	"jewelry.abgdrv.com/internal/validator"
	"strings"
	"time"
)

type Watch struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Brand     string    `json:"brand,omitempty"`
	Model     string    `json:"model,omitempty"`
	DialColor string    `json:"dial_color"`
	StrapType string    `json:"strap_type"`
	Diameter  int64     `json:"diameter"`
	Energy    string    `json:"energy"`
	Gender    string    `json:"gender"`
	Price     float64   `json:"price"`
	ImageURL  string    `json:"image_url"`
}

func ValidateWatch(v *validator.Validator, watch *Watch) {
	v.Check(watch.Brand != "", "brand", "must be provided")
	v.Check(len(watch.Brand) <= 500, "brand", "must not be more than 500 bytes long")

	v.Check(watch.Model != "", "model", "must be provided")
	v.Check(len(watch.Model) <= 500, "model", "must not be more than 500 bytes long")

	v.Check(watch.DialColor != "", "dial_color", "must be provided")
	v.Check(len(watch.DialColor) <= 500, "dial_color", "must not be more than 500 bytes long")

	v.Check(watch.StrapType != "", "strap_type", "must be provided")
	v.Check(len(watch.StrapType) <= 500, "strap_type", "must not be more than 500 bytes long")

	//v.Check(strings.ToLower(watch.Gender) == "male" && watch.Diameter >= 38 && watch.Diameter <= 46, "diameter", "for men should be between 38mm and 46mm")
	//v.Check(strings.ToLower(watch.Gender) == "female" && watch.Diameter >= 26 && watch.Diameter <= 36, "diameter", "for women should be between 26mm and 36mm")

	v.Check(strings.ToLower(watch.Energy) == "mechanical" || strings.ToLower(watch.Energy) == "quartz", "energy", "fall into two main categories: Mechanical and Quartz")

	v.Check(strings.ToLower(watch.Gender) == "male" || strings.ToLower(watch.Gender) == "female", "gender", "Gender can be male or female")

	v.Check(watch.Price > 0, "price", "can not be equal or less than 0")

	v.Check(watch.ImageURL != "", "image_url", "must be provided")
}

type WatchModel struct {
	DB *sql.DB
}

func (w WatchModel) Insert(watch *Watch) error {
	query := `INSERT INTO watches (brand, model, dial_color, strap_type, diameter, energy, gender, price, image_url) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
				RETURNING id, created_at`

	args := []interface{}{
		watch.Brand,
		watch.Model,
		watch.DialColor,
		watch.StrapType,
		watch.Diameter,
		watch.Energy,
		watch.Gender,
		watch.Price,
		watch.ImageURL,
	}

	return w.DB.QueryRow(query, args...).Scan(&watch.ID, &watch.CreatedAt)
}

func (w WatchModel) Get(id int64) (*Watch, error) {
	return nil, nil
}

func (w WatchModel) Update(watch *Watch) error {
	return nil
}

func (w WatchModel) Delete(id int64) error {
	return nil
}

type MockWatchModel struct{}

func (m MockWatchModel) Insert(watch *Watch) error {
	return nil
}

func (m MockWatchModel) Get(id int64) (*Watch, error) {
	return nil, nil
}

func (m MockWatchModel) Update(watch *Watch) error {
	return nil
}

func (m MockWatchModel) Delete(id int64) error {
	return nil
}
