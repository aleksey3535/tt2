package repository

import (
	"database/sql"
	"fmt"
	"time"
	"tt2/internal/models"

	"github.com/google/uuid"
)

func (r *Repository) Create(sub models.SubscriptionForCreate) (int, error) {
	var id int
	query := `INSERT INTO subscriptions (service_name, price, user_id, start_date) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(query, sub.ServiceName, sub.Price, sub.UserID, parseToDate(sub.StartDate)).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) GetAll() ([]models.Subscription, error) {
	var subs []models.Subscription
	query := `SELECT id, service_name, price, user_id, start_date FROM subscriptions`
	err := r.db.Select(&subs, query)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (r *Repository) GetByID(id int) (models.Subscription, error) {
	var sub models.Subscription
	query := `SELECT id, service_name, price, user_id, start_date FROM subscriptions WHERE id=$1`
	err := r.db.Get(&sub, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Subscription{}, ErrNotFound
		}
		return models.Subscription{}, err
	}
	return sub, nil
}

func (r *Repository) Update(id int, sub models.SubscriptionForCreate) error {
	query := `UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, start_date=$4 WHERE id=$5`
	_, err := r.db.Exec(query, sub.ServiceName, sub.Price, sub.UserID, parseToDate(sub.StartDate), id)
	return err
}
func (r *Repository) Delete(id int) error {
	query := `DELETE FROM subscriptions WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) Total(userID *uuid.UUID, serviceName string, startDate *time.Time) (int, error) {
	fmt.Println(userID, serviceName, startDate)
	var total int
	query := `SELECT COALESCE(SUM(price), 0) FROM subscriptions
				WHERE ($1::uuid IS NULL OR user_id = $1::uuid)
  				AND ($2 = '' OR service_name = $2)
  				AND ($3::date IS NULL OR start_date >= $3::date)
`
	err := r.db.Get(&total, query, userID, serviceName, startDate)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func parseToDate(startDate string) time.Time {
	t, _ := time.Parse("01-2006", startDate)
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
}
