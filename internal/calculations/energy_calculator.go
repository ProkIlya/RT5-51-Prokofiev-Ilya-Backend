// internal/calculations/energy_calculator.go
package calculations

import "tesla-app/internal/models"

const (
	MassKg         = 2068
	AirDensityKgM3 = 1.225
	FrontalAreaM2  = 2.34
	GravityMpS2    = 9.8
)

func CalculateRemainingCharge(trip *models.TripCalculationCharge) float64 {
	totalEnergy := 0.0

	for _, tripScenario := range trip.Scenarios {
		scenario := tripScenario.Scenario
		value := tripScenario.Value

		if scenario.Type == models.RoadType {
			// Расчет энергии на движение для дорожных условий
			F_rolling := float64(MassKg) * GravityMpS2 * scenario.RollingCoeff
			speedMs := scenario.Speed / 3.6
			F_air := 0.5 * AirDensityKgM3 * scenario.AeroCoeff * FrontalAreaM2 * speedMs * speedMs
			E_drive := ((F_rolling + F_air) * value * 1000) / 3600000
			totalEnergy += E_drive
		} else if scenario.Type == models.ComfortType {
			// Расчет энергии на системы комфорта
			E_systems := scenario.SystemConsuption * value
			totalEnergy += E_systems
		}
	}

	remainingCharge := trip.StartCharge - totalEnergy

	if remainingCharge < 0 {
		return 0
	}
	return remainingCharge
}
