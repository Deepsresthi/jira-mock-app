package controllers

import (
	"time"

	"jira-mock-app/app"
	"jira-mock-app/app/middleware"
	"jira-mock-app/app/models"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type Sprints struct {
	*revel.Controller
	DB *gorm.DB
}

func (c Sprints) Index() revel.Result {
	return c.RenderText("Sprints Index page")
}

func (c Sprints) List() revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "List")
	if err != nil {
		return c.Forbidden("Access denied")
	}

	db := app.DB
	var sprints []models.Sprints
	if err := db.Find(&sprints).Error; err != nil {
		revel.AppLog.Errorf("Error fetching sprints: %v", err)
		return c.RenderError(err)
	}

	// Check if there are sprints
	if len(sprints) == 0 {
		return c.RenderJSON(map[string]interface{}{
			"message": "No sprints",
		})
	}

	return c.RenderJSON(sprints)
}

func (c Sprints) New() revel.Result {
	return c.Render()
}

func (c Sprints) Create(sprintId uint, name string, startDate, endDate time.Time) revel.Result {
	// Apply JWT authorization middleware to restrict access
	//print
	err := middleware.TokenAuth(c.Controller, "Create")
	if err != nil {
		return c.Forbidden("Access denied")
	}

	db := app.DB
	newSprint := models.Sprints{
		SprintID:  sprintId,
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
	}

	if err := db.Create(&newSprint).Error; err != nil {
		revel.AppLog.Errorf("Error creating sprint: %v", err)
		return c.RenderError(err)
	}

	return c.Redirect(Sprints.List)
}

func (c Sprints) Show(sprintId uint) revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "Show")
	if err != nil {
		return c.Forbidden("Access denied")
	}

	db := app.DB
	var sprint models.Sprints
	if err := db.Where("sprint_id = ?", sprintId).First(&sprint).Error; err != nil {
		revel.AppLog.Errorf("Error fetching sprint: %v", err)
		return c.NotFound("Sprint not found")
	}

	return c.Render(sprint)
}

func (c Sprints) Edit(sprintId uint) revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "Edit")
	if err != nil {
		return c.Forbidden("Access denied")
	}

	db := app.DB
	var sprint models.Sprints
	if err := db.Where("sprint_id = ?", sprintId).First(&sprint).Error; err != nil {
		revel.AppLog.Errorf("Error fetching sprint: %v", err)
		return c.NotFound("Sprint not found")
	}

	return c.Render(sprint)
}

func (c Sprints) Update(sprintId uint, name string, startDate, endDate time.Time) revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "Update")
	if err != nil {
		return c.Forbidden("Access denied")
	}

	db := app.DB
	var sprint models.Sprints
	if err := db.Where("sprint_id = ?", sprintId).First(&sprint).Error; err != nil {
		revel.AppLog.Errorf("Error fetching sprint: %v", err)
		return c.NotFound("Sprint not found")
	}

	sprint.Name = name
	sprint.StartDate = startDate
	sprint.EndDate = endDate

	if err := db.Save(&sprint).Error; err != nil {
		revel.AppLog.Errorf("Error updating sprint: %v", err)
	}

	return c.Redirect(Sprints.List)
}

func (c Sprints) Delete(sprintId uint) revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "Delete")
	if err != nil {
		return c.Forbidden("Access denied")
	}

	db := app.DB
	var sprint models.Sprints
	if err := db.Where("sprint_id = ?", sprintId).First(&sprint).Error; err != nil {
		revel.AppLog.Errorf("Error fetching sprint: %v", err)
		return c.NotFound("Sprint not found")
	}

	if err := db.Delete(&sprint).Error; err != nil {
		revel.AppLog.Errorf("Error deleting sprint: %v", err)
	}

	return c.Redirect(Sprints.List)
}
