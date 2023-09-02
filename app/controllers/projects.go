package controllers

import (
	"jira-mock-app/app"
	"jira-mock-app/app/middleware"
	"jira-mock-app/app/models"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type Projects struct {
	*revel.Controller
	DB *gorm.DB
}

func (c Projects) Index() revel.Result {
	return c.RenderJSON(map[string]interface{}{
		"message": "Projects Index page",
	})
}

func (c Projects) List() revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "List")
	if err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Access denied",
		})
	}

	db := app.DB
	var projects []models.Project
	if err := db.Find(&projects).Error; err != nil {
		revel.AppLog.Errorf("Error fetching projects: %v", err)
		return c.RenderJSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Check if there are projects
	if len(projects) == 0 {
		return c.RenderJSON(map[string]interface{}{
			"message": "No projects",
		})
	}

	return c.RenderJSON(projects)
}

func (c Projects) New() revel.Result {
	return c.RenderJSON(map[string]interface{}{})
}

func (c Projects) Create(projectID uint, projectName, teamName, projectOwner string) revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "Create")
	if err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Access denied",
		})
	}

	db := app.DB
	newProject := models.Project{
		ProjectID:    projectID,
		ProjectName:  projectName,
		TeamName:     teamName,
		ProjectOwner: projectOwner,
	}

	if err := db.Create(&newProject).Error; err != nil {
		revel.AppLog.Errorf("Error creating project: %v", err)
		return c.RenderJSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.RenderJSON(map[string]interface{}{
		"message": "Project created successfully",
	})
}

func (c Projects) Show(projectID uint) revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "Show")
	if err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Access denied",
		})
	}

	db := app.DB
	var project models.Project
	if err := db.Where("project_id = ?", projectID).First(&project).Error; err != nil {
		revel.AppLog.Errorf("Error fetching project: %v", err)
		return c.RenderJSON(map[string]interface{}{
			"error": "Project not found",
		})
	}

	return c.RenderJSON(project)
}

func (c Projects) Edit(projectID uint) revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "Edit")
	if err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Access denied",
		})
	}

	db := app.DB
	var project models.Project
	if err := db.Where("project_id = ?", projectID).First(&project).Error; err != nil {
		revel.AppLog.Errorf("Error fetching project: %v", err)
		return c.RenderJSON(map[string]interface{}{
			"error": "Project not found",
		})
	}

	return c.RenderJSON(project)
}

func (c Projects) Update(projectID uint, projectName, teamName, projectOwner string) revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "Update")
	if err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Access denied",
		})
	}

	db := app.DB
	var project models.Project
	if err := db.Where("project_id = ?", projectID).First(&project).Error; err != nil {
		revel.AppLog.Errorf("Error fetching project: %v", err)
		return c.RenderJSON(map[string]interface{}{
			"error": "Project not found",
		})
	}

	project.ProjectName = projectName
	project.TeamName = teamName
	project.ProjectOwner = projectOwner

	if err := db.Save(&project).Error; err != nil {
		revel.AppLog.Errorf("Error updating project: %v", err)
		return c.RenderJSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.RenderJSON(map[string]interface{}{
		"message": "Project updated successfully",
	})
}

func (c Projects) Delete(projectID uint) revel.Result {
	// Apply JWT authorization middleware to restrict access
	err := middleware.TokenAuth(c.Controller, "Delete")
	if err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Access denied",
		})
	}

	db := app.DB
	var project models.Project
	if err := db.Where("project_id = ?", projectID).First(&project).Error; err != nil {
		revel.AppLog.Errorf("Error fetching project: %v", err)
		return c.RenderJSON(map[string]interface{}{
			"error": "Project not found",
		})
	}

	if err := db.Delete(&project).Error; err != nil {
		revel.AppLog.Errorf("Error deleting project: %v", err)
		return c.RenderJSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.RenderJSON(map[string]interface{}{
		"message": "Project deleted successfully",
	})
}
