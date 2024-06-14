package routes

import (
	"github.com/PhilanderNews/BackendGin/controllers"
	"github.com/gin-gonic/gin"
)

func BeritaRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/TambahBerita", controllers.TambahBerita)
	incomingRoutes.POST("/AmbilSatuBerita", controllers.AmbilSatuBerita)
	incomingRoutes.GET("/AmbilSemuaBerita", controllers.AmbilSemuaBerita)
	incomingRoutes.PUT("/EditBerita", controllers.EditBerita)
	incomingRoutes.DELETE("/HapusBerita", controllers.HapusBerita)
}
