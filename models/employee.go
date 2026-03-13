package models

import (
	"time"
)

type Employee struct {
	ID                   int       `json:"id"`
	UniqueKey            string    `json:"uniqueKey"`
	Name                 string    `json:"name"`
	IdentificationNumber string    `json:"identificationNumber"`
	Age                  int       `json:"age"`
	Address              string    `json:"address"`
	Occupation           string    `json:"occupation"`
	PlaceOfBirth         string    `json:"placeOfBirth"`
	DateOfBirth          string    `json:"dateOfBirth"`
	CreatedAt            time.Time `json:"createdAt"`
}

type EmployeeResponse struct {
	ID                   int    `json:"id"`
	UniqueKey            string `json:"uniqueKey"`
	Name                 string `json:"name"`
	IdentificationNumber string `json:"identificationNumber"`
	Age                  int    `json:"age"`
	Address              string `json:"address"`
	Occupation           string `json:"occupation"`
	PlaceOfBirth         string `json:"placeOfBirth"`
	DateOfBirth          string `json:"dateOfBirth"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func CountAge(str string) (int, error) {
	dob, err := time.Parse("2006-01-02", str)
	if err != nil {
		return 0, err
	}

	today := time.Now()
	age := today.Year() - dob.Year()

	if today.YearDay() < dob.YearDay() {
		age--
	}

	return age, nil
}
