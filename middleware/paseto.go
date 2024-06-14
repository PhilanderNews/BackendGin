package middleware

import (
	"net/http"
	"os"

	"github.com/PhilanderNews/BackendGin/helpers"
	"github.com/PhilanderNews/BackendGin/models"
	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Header login tidak ditemukan"})
		c.Abort()
		return
	}
	tokenname := helpers.DecodeGetName(os.Getenv("publickey"), header)
	tokenusername := helpers.DecodeGetUsername(os.Getenv("publickey"), header)
	tokenrole := helpers.DecodeGetRole(os.Getenv("publickey"), header)
	if tokenusername == "" || tokenrole == "" {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Belum Login"})
		c.Abort()
		return
	}
	c.Set("name", tokenname)
	c.Set("username", tokenusername)
	c.Set("role", tokenrole)
}
