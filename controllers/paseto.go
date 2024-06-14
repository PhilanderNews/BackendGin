package controllers

import (
	"net/http"

	"github.com/PhilanderNews/BackendGin/middleware"
	"github.com/PhilanderNews/BackendGin/models"
	"github.com/PhilanderNews/BackendGin/utils"
	"github.com/gin-gonic/gin"
)

func Authorization(publickey, mongoenv, dbname, collname string) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func TokenValue(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var response models.CredentialUser
	// Authorization
	middleware.Authorization(c)
	name := c.GetString("name")
	username := c.GetString("username")
	role := c.GetString("role")
	// Cek Username
	if !utils.UsernameExists(mconn, "users", models.Users{Username: username}) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Akun tidak ditemukan"})
		c.Abort()
		return
	}
	// Create Response
	response.Status = true
	response.Message = "Berhasil decode token"
	response.Data.Name = name
	response.Data.Username = username
	response.Data.Role = role

	c.JSON(http.StatusOK, response)
}
