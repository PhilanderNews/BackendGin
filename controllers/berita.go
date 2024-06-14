package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/PhilanderNews/BackendGin/middleware"
	"github.com/PhilanderNews/BackendGin/models"
	"github.com/PhilanderNews/BackendGin/utils"
	"github.com/gin-gonic/gin"
)

func TambahBerita(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var berita models.Berita
	err := c.BindJSON(&berita)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	//Define waktu
	wib, timeErr := time.LoadLocation("Asia/Jakarta")
	if timeErr != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing time location: " + err.Error()})
		return
	}
	currentTime := time.Now().In(wib)
	timeStringBerita := currentTime.Format("Monday, 2 January 2006 15:04 MST")
	// Authorization
	middleware.Authorization(c)
	if c.IsAborted() {
		return
	}
	role := c.GetString("role")
	name := c.GetString("name")
	// Cek role
	if role != "admin" && role != "author" {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
		c.Abort()
		return
	}
	// Cek apakah id berita telah ada
	if utils.IDBeritaExists(mconn, "berita", berita) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "ID telah ada"})
		return
	}
	// Ensure all paragraphs have at least a space
	for i := 1; i <= 10; i++ {
		field := reflect.ValueOf(&berita.Konten).Elem().FieldByName(fmt.Sprintf("Paragraf%d", i))
		if field.String() == "" {
			field.SetString(" ")
		}
	}
	// Insert data berita
	berita.Penulis = name
	berita.Waktu = timeStringBerita
	utils.InsertBerita(mconn, "berita", berita)
	c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil input data"})
}

func AmbilSatuBerita(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var berita models.Berita
	err := c.BindJSON(&berita)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	if berita.ID == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID"})
		return
	}
	// Cek apakah id berita ada
	if utils.IDBeritaExists(mconn, "berita", berita) {
		// Get data berita
		databerita := utils.FindBerita(mconn, "berita", berita)
		c.JSON(http.StatusOK, databerita)
	} else {
		c.JSON(http.StatusOK, models.Pesan{Status: false, Message: "Berita tidak ditemukan"})
	}
}

func AmbilSemuaBerita(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	// Get data berita
	databerita, err := utils.GetAllBerita(mconn, "berita")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetAllDoc error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, databerita)
}

func EditBerita(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var berita models.Berita
	err := c.BindJSON(&berita)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	// Authorization
	middleware.Authorization(c)
	if c.IsAborted() {
		return
	}
	role := c.GetString("role")
	name := c.GetString("name")
	// Cek role
	namapenulis := utils.FindBerita(mconn, "berita", berita)
	if !(role == "admin" || name == namapenulis.Penulis) {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
		c.Abort()
		return
	}
	// Cek apakah berita ID ada
	if berita.ID == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID"})
		return
	}
	if utils.IDBeritaExists(mconn, "berita", berita) {
		for i := 1; i <= 10; i++ {
			field := reflect.ValueOf(&berita.Konten).Elem().FieldByName(fmt.Sprintf("Paragraf%d", i))
			if field.String() == "" {
				field.SetString(" ")
			}
		}
		// Update data berita
		utils.UpdateBerita(mconn, "berita", berita)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil update " + berita.ID + " dari database"})
	} else {
		c.JSON(http.StatusOK, models.Pesan{Status: false, Message: "Berita tidak ditemukan"})
	}
}

func HapusBerita(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var berita models.Berita
	err := c.BindJSON(&berita)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	// Authorization
	middleware.Authorization(c)
	if c.IsAborted() {
		return
	}
	role := c.GetString("role")
	name := c.GetString("name")
	// Cek role
	namapenulis := utils.FindBerita(mconn, "berita", berita)
	if !(role == "admin" || name == namapenulis.Penulis) {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
		c.Abort()
		return
	}
	// Cek apakah berita ID ada
	if berita.ID == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID"})
		return
	}
	if utils.IDBeritaExists(mconn, "berita", berita) {
		// Delete data berita
		utils.DeleteBerita(mconn, "berita", berita)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil delete " + berita.ID + " dari database"})
	} else {
		c.JSON(http.StatusOK, models.Pesan{Status: false, Message: "Berita tidak ditemukan"})
	}
}
