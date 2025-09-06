// internal/handlers/handlers.go
package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"tesla-app/internal/calculations"
	"tesla-app/internal/models"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	searchQuery := c.Query("search")
	var filteredScenarios []*models.DrivingScenario

	if searchQuery != "" {
		for _, scenario := range models.DrivingScenarios {
			if strings.Contains(strings.ToLower(scenario.Name), strings.ToLower(searchQuery)) ||
				strings.Contains(strings.ToLower(scenario.Description), strings.ToLower(searchQuery)) {
				filteredScenarios = append(filteredScenarios, scenario)
			}
		}
	} else {
		filteredScenarios = models.DrivingScenarios
	}

	trip := models.Trips[1]

	c.HTML(http.StatusOK, "drive_menu.html", gin.H{
		"Scenarios":   filteredScenarios,
		"SearchQuery": searchQuery,
		"Trip":        trip,
	})
}

func ScenarioHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var scenario *models.DrivingScenario
	for _, s := range models.DrivingScenarios {
		if s.ID == id {
			scenario = s
			break
		}
	}

	if scenario == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	trip := models.Trips[1] // Добавляем получение заявки

	c.HTML(http.StatusOK, "scenario.html", gin.H{
		"Scenario": scenario,
		"Trip":     trip, // Добавляем заявку в данные для шаблона
	})
}

func TripHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	trip, exists := models.Trips[id]
	if !exists {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	remainingCharge := calculations.CalculateRemainingCharge(trip)

	c.HTML(http.StatusOK, "trip_calculation.html", gin.H{
		"Trip":            trip,
		"RemainingCharge": remainingCharge,
	})
}
