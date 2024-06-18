package main

import (
	"fmt"
	"net/http"

	"github.com/PhilanderNews/BackendGin/models"
	"github.com/PhilanderNews/BackendGin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error occurred on .env file, please check ", err.Error())
	}
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"0.0.0.0"})

	routes.UserRoutes(router)
	routes.BeritaRoutes(router)
	routes.KomentarRoutes(router)
	routes.VerifikasiRoutes(router)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, models.Pesan{Status: false, Message: "Page not found"})
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
