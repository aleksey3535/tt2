package models

import (
	"time"
)


// SubscriptionForCreate описывает тело запроса создания подписки
//
// swagger:model SubscriptionForCreate
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
