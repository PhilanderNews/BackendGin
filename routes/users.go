package routes

import (
	"github.com/PhilanderNews/BackendGin/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/TokenValue", controllers.TokenValue)
	incomingRoutes.POST("/Registrasi", controllers.Registrasi)
	incomingRoutes.POST("/Login", controllers.Login)
	incomingRoutes.POST("/AmbilSatuUser", controllers.AmbilSatuUser)
	incomingRoutes.GET("/AmbilSemuaUser", controllers.AmbilSemuaUser)
	incomingRoutes.PUT("/EditUser", controllers.EditUser)
	incomingRoutes.DELETE("/HapusUser", controllers.HapusUser)
}
