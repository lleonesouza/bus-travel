package dtos

import (
	"time"

	"github.com/go-playground/validator"
)

type BusTravel struct {
	ID                uint      `json:"id" example:"1"`
	Origin            string    `json:"origin" validate:"required" example:"São Paulo"`
	Destination       string    `json:"destination,omitempty" example:"Rio de Janeiro"`
	DepartureTime     time.Time `json:"departureTime" validate:"required" example:"2022-01-01T10:00:00Z"`
	ArrivalTime       time.Time `json:"arrivalTime" validate:"required" example:"2022-01-01T15:00:00Z"`
	Price             float64   `json:"price" validate:"required,gte=0" example:"50.00"`
	HasWiFi           bool      `json:"hasWiFi" example:"true"`
	HasBathroom       bool      `json:"hasBathroom" example:"true"`
	HasAirConditioner bool      `json:"hasAirConditioner" example:"true"`
}

// Create Input
type BusTravelCreateInput struct {
	Origin            string    `json:"origin" validate:"required" example:"São Paulo"`
	Destination       string    `json:"destination,omitempty" example:"Rio de Janeiro"`
	DepartureTime     time.Time `json:"departureTime" validate:"required" example:"2022-01-01T10:00:00Z"`
	ArrivalTime       time.Time `json:"arrivalTime" validate:"required" example:"2022-01-01T15:00:00Z"`
	Price             float64   `json:"price" validate:"required,gte=0" example:"50.00"`
	HasWiFi           bool      `json:"hasWiFi" example:"true"`
	HasBathroom       bool      `json:"hasBathroom" example:"true"`
	HasAirConditioner bool      `json:"hasAirConditioner" example:"true"`
}

func (i *BusTravelCreateInput) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}

// Update Input
type BusTravelUpdateInput struct {
	Origin            *string    `json:"origin,omitempty" validate:"omitempty"`
	Destination       *string    `json:"destination,omitempty"`
	DepartureTime     *time.Time `json:"departureTime,omitempty" example:"2022-01-01T10:00:00Z"`
	ArrivalTime       *time.Time `json:"arrivalTime,omitempty" example:"2022-01-01T15:00:00Z"`
	Price             *float64   `json:"price,omitempty" example:"50.00"`
	HasWiFi           *bool      `json:"hasWiFi,omitempty" example:"true"`
	HasBathroom       *bool      `json:"hasBathroom,omitempty" example:"true"`
	HasAirConditioner *bool      `json:"hasAirConditioner,omitempty" example:"true"`
}

func (i *BusTravelUpdateInput) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}
