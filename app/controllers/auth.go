package controllers

import (
	"jira-mock-app/app"
	"jira-mock-app/app/models"

	"jira-mock-app/app/middleware"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Login() revel.Result {
	return c.Render()
}

func (c Auth) DoLogin(username, password string) revel.Result {
	db := app.DB // Directly access the DB instance
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "User not found",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Incorrect password",
		})
	}

	// Generate JWT token
	tokenString, err := middleware.GenerateJWTToken(username)
	if err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Token generation failed",
		})
	}

	// Include the token in the response
	return c.RenderJSON(map[string]interface{}{
		"token": tokenString,
	})
}

func (c Auth) Signup() revel.Result {
	return c.Render()
}

func (c Auth) DoSignup(username, mailID, password string) revel.Result {
	// var requestBody struct {
	// 	Username     string `json:"username"`
	// 	MailID       string `json:"mailID"`
	// 	PasswordHash string `json:"password"`
	// }

	// err := c.Params.BindJSON(&requestBody)
	// if err != nil {
	// 	return c.RenderJSON(map[string]interface{}{
	// 		"error": "Invalid JSON format",
	// 	})
	// }

	db := app.DB // Directly access the DB instance
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Signup failed",
		})
	}

	newUser := models.User{
		Username:     username,
		MailID:       mailID,
		PasswordHash: string(hashedPassword),
	}
	if err := db.Create(&newUser).Error; err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Signup failed",
		})
	}
	return c.RenderJSON(map[string]interface{}{
		"message": "Signup successful",
	})
}
