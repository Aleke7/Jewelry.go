package data

import (
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
}

func ValidateWatch(v *validator.Validator, watch Watch) {
	v.Check(watch.Brand != "", "brand", "must be provided")
	v.Check(len(watch.Brand) <= 500, "brand", "must not be more than 500 bytes long")

	v.Check(watch.Model != "", "model", "must be provided")
	v.Check(len(watch.Model) <= 500, "model", "must not be more than 500 bytes long")

	v.Check(watch.DialColor != "", "dial_color", "must be provided")
	v.Check(len(watch.DialColor) <= 500, "dial_color", "must not be more than 500 bytes long")

	v.Check(watch.StrapType != "", "strap_type", "must be provided")
	v.Check(len(watch.StrapType) <= 500, "strap_type", "must not be more than 500 bytes long")

	v.Check(strings.ToLower(watch.Gender) == "male" && watch.Diameter >= 38 && watch.Diameter <= 46, "diameter", "for men should be between 38mm and 46mm")
	v.Check(strings.ToLower(watch.Gender) == "female" && watch.Diameter >= 26 && watch.Diameter <= 36, "diameter", "for women should be between 26mm and 36mm")

	v.Check(strings.ToLower(watch.Energy) == "mechanical" || strings.ToLower(watch.Energy) == "quartz", "energy", "fall into two main categories: Mechanical and Quartz")

	v.Check(strings.ToLower(watch.Gender) == "male" || strings.ToLower(watch.Gender) == "female", "gender", "Gender can be male or female")

	v.Check(watch.Price > 0, "price", "can not be equal or less than 0")
}
