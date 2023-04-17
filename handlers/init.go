package handlers

import "github.com/lleonesouza/bus-travel/services"

type Handlers struct {
	BusTravelHandler *BusTravelHandler
	Health           *HealthHandler
}

func MakeHandlers(service *services.BusTravelService) *Handlers {

	return &Handlers{
		BusTravelHandler: &BusTravelHandler{service},
		Health:           &HealthHandler{},
	}

}
