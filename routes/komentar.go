package routes

import (
	"github.com/PhilanderNews/BackendGin/controllers"
	"github.com/gin-gonic/gin"
)

func KomentarRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/TambahKomentar", controllers.TambahKomentar)
	incomingRoutes.POST("/AmbilSatuKomentar", controllers.AmbilSatuKomentar)
	incomingRoutes.GET("/AmbilSemuaKomentar", controllers.AmbilSemuaKomentar)
	incomingRoutes.PUT("/EditKomentar", controllers.EditKomentar)
	incomingRoutes.DELETE("/HapusKomentar", controllers.HapusKomentar)
}
