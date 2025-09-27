package handler

import (
	"time"

	"github.com/google/uuid"
)

func validateUserID(userID string) (*uuid.UUID, error) {
	var zeroValue *uuid.UUID = nil
	if userID == "" {
		return zeroValue, nil
	}
	id, err := uuid.Parse(userID)
	if err != nil {
		return zeroValue, err
	}
	return &id, nil
}

func validateStartDate(startDate string) (*time.Time, error) {
	var zeroValue *time.Time = nil
	if startDate == "" {
		return zeroValue, nil
	}
	t, err := time.Parse("01-2006", startDate)
	if err != nil {
		return zeroValue, err
	}
	curTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	return &curTime, nil
}


