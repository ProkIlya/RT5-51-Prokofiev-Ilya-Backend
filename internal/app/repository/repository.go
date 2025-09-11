package repository

import "tesla-app/internal/models"

type Repository struct {
	DrivingScenarios []*models.DrivingScenario
	Trips            map[int]*models.TripCalculationCharge
}

func NewRepository() (*Repository, error) {
	return &Repository{
		DrivingScenarios: models.DrivingScenarios,
		Trips:            models.Trips,
	}, nil
}

func (r *Repository) GetScenarios() []*models.DrivingScenario {
	return r.DrivingScenarios
}

func (r *Repository) GetScenarioByID(id int) *models.DrivingScenario {
	for _, scenario := range r.DrivingScenarios {
		if scenario.ID == id {
			return scenario
		}
	}
	return nil
}

func (r *Repository) GetTripByID(id int) *models.TripCalculationCharge {
	return r.Trips[id]
}
