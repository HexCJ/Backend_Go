package handlers

import (
	"net/http"
	"strconv"

	"gin-user-api/internal/models"
	"gin-user-api/internal/services"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	Service services.ProfileService
}

// POST /api/users/:id/profile
func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	var profile models.Profile
	c.ShouldBindJSON(&profile)
	profile.UserID = uint(userID)

	h.Service.Create(&profile)

	c.JSON(http.StatusCreated, profile)
}

// PUT /api/users/:id/profile
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	var profile models.Profile
	h.Service.GetByUserID(&profile, uint(userID))

	c.ShouldBindJSON(&profile)
	h.Service.Update(&profile)

	c.JSON(http.StatusOK, profile)
}
