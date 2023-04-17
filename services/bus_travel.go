package services

import (
	"context"

	"github.com/lleonesouza/bus-travel/modules/dtos"
	"github.com/lleonesouza/bus-travel/prisma/db"
)

type BusTravelService struct {
	client *db.PrismaClient
}

func NewBusTravelService(client *db.PrismaClient) *BusTravelService {
	return &BusTravelService{client: client}
}

func (s *BusTravelService) Create(ctx context.Context, input dtos.BusTravelCreateInput) (*db.BusTravelModel, error) {
	createdBusTravel, err := s.client.BusTravel.CreateOne(
		db.BusTravel.Origin.Set(input.Origin),
		db.BusTravel.DepartureTime.Set(input.DepartureTime),
		db.BusTravel.ArrivalTime.Set(input.ArrivalTime),
		db.BusTravel.Price.Set(input.Price),
		db.BusTravel.Destination.Set(input.Destination),
		db.BusTravel.HasWiFi.Set(input.HasWiFi),
		db.BusTravel.HasBathroom.Set(input.HasBathroom),
		db.BusTravel.HasAirConditioner.Set(input.HasAirConditioner),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdBusTravel, nil
}

func (s *BusTravelService) Update(ctx context.Context, id int, input dtos.BusTravelUpdateInput) (*db.BusTravelModel, error) {
	updatedBusTravel, err := s.client.BusTravel.FindUnique(
		db.BusTravel.ID.Equals(id),
	).Update(
		db.BusTravel.Origin.SetIfPresent(input.Origin),
		db.BusTravel.Destination.SetIfPresent(input.Destination),
		db.BusTravel.DepartureTime.SetIfPresent(input.DepartureTime),
		db.BusTravel.ArrivalTime.SetIfPresent(input.ArrivalTime),
		db.BusTravel.Price.SetIfPresent(input.Price),
		db.BusTravel.HasWiFi.SetIfPresent(input.HasWiFi),
		db.BusTravel.HasBathroom.SetIfPresent(input.HasBathroom),
		db.BusTravel.HasAirConditioner.SetIfPresent(input.HasAirConditioner),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updatedBusTravel, nil
}

func (s *BusTravelService) Get(ctx context.Context, id int) (*db.BusTravelModel, error) {
	busTravel, err := s.client.BusTravel.FindUnique(
		db.BusTravel.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return busTravel, nil
}

func (s *BusTravelService) GetAll(ctx context.Context) ([]db.BusTravelModel, error) {
	busTravels, err := s.client.BusTravel.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return busTravels, nil
}

func (s *BusTravelService) Delete(ctx context.Context, id int) error {
	_, err := s.client.BusTravel.FindUnique(
		db.BusTravel.ID.Equals(id),
	).Delete().Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
