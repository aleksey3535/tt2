package models

import (
	"time"
)

// type CustomDate time.Time

// func (cd *CustomDate) UnmarshalJSON(b []byte) error {
//     s := strings.Trim(string(b), "\"")
//     t, err := time.Parse("01-2006", s)
//     if err != nil {
//         return err
//     }
//     *cd = CustomDate(t)
//     return nil
// }

type SubscriptionForCreate struct {
	ServiceName string `json:"service_name" db:"service_name"`
	Price       int    `json:"price" db:"price"`
	UserID      string `json:"user_id" db:"user_id"`
	StartDate   string `json:"start_date" db:"start_date"`
}

type Subscription struct {
	Id          int       `json:"id" db:"id"`
	ServiceName string    `json:"service_name" db:"service_name"`
	Price       int       `json:"price" db:"price"`
	UserID      string    `json:"user_id" db:"user_id"`
	StartDate   time.Time `json:"start_date" db:"start_date"`
}
