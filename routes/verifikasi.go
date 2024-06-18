package routes

import (
	"github.com/PhilanderNews/BackendGin/controllers"
	"github.com/gin-gonic/gin"
)

func VerifikasiRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/GenerateVerifikasi", controllers.GenerateVerifikasi)
}
