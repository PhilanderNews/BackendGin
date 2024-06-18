package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/PhilanderNews/BackendGin/helpers"
	"github.com/PhilanderNews/BackendGin/middleware"
	"github.com/PhilanderNews/BackendGin/models"
	"github.com/PhilanderNews/BackendGin/utils"
	"github.com/gin-gonic/gin"
)

func TambahKomentar(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var komentar models.Komentar
	var berita models.Berita
	err := c.BindJSON(&komentar)
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
	timeStringKomentar := currentTime.Format("January 2, 2006")
	// Cek role
	header := c.Request.Header.Get("Authorization")
	role := helpers.DecodeGetRole(os.Getenv("publickey"), header)
	username := helpers.DecodeGetUsername(os.Getenv("publickey"), header)
	if role != "admin" && role != "author" && role != "user" {
		// Cek apakah id komentar telah ada
		if komentar.ID == "" {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID"})
			return
		}
		if utils.IDKomentarExists(mconn, "komentar", komentar) {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "ID telah ada"})
			return
		}
		// Cek apakah id berita ada
		if komentar.ID_berita == "" {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID Berita"})
			return
		}
		berita.ID = komentar.ID_berita
		if !utils.IDBeritaExists(mconn, "berita", berita) {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Berita tidak ditemukan"})
			return
		}
		komentar.Username = "Anonymous"
		komentar.Tanggal = timeStringKomentar
		utils.InsertKomentar(mconn, "komentar", komentar)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil Input data tanpa login"})
	}
	// Cek apakah id komentar telah ada
	if komentar.ID == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID"})
		return
	}
	if utils.IDKomentarExists(mconn, "komentar", komentar) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "ID telah ada"})
		return
	}
	// Cek apakah id berita ada
	if komentar.ID_berita == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID Berita"})
		return
	}
	berita.ID = komentar.ID_berita
	if !utils.IDBeritaExists(mconn, "berita", berita) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Berita tidak ditemukan"})
		return
	}
	// Insert data komentar
	komentar.Username = username
	komentar.Tanggal = timeStringKomentar
	utils.InsertKomentar(mconn, "komentar", komentar)
	c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil input data"})
}

func AmbilSatuKomentar(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var komentar models.Komentar
	err := c.BindJSON(&komentar)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	if komentar.ID == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID"})
		return
	}
	// Cek apakah id komentar ada
	if utils.IDKomentarExists(mconn, "komentar", komentar) {
		// Get data komentar
		datakomentar := utils.FindKomentar(mconn, "komentar", komentar)
		c.JSON(http.StatusOK, datakomentar)
	} else {
		c.JSON(http.StatusOK, models.Pesan{Status: false, Message: "Komentar tidak ditemukan"})
	}
}

func AmbilSemuaKomentar(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	// Get data user
	datakomentar, err := utils.GetAllKomentar(mconn, "komentar")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetAllDoc error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, datakomentar)
}

func EditKomentar(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var komentar models.Komentar
	err := c.BindJSON(&komentar)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	// Authorization
	middleware.Authorization(c)
	if c.IsAborted() {
		return
	}
	username := c.GetString("username")
	// Cek role
	namakomentator := utils.FindKomentar(mconn, "komentar", komentar)
	if username != namakomentator.Username {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
		c.Abort()
		return
	}
	// Cek apakah komentar ID ada
	if komentar.ID == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID"})
		return
	}
	if utils.IDKomentarExists(mconn, "komentar", komentar) {
		// Update data komentar
		komentar.Username = username
		komentar.ID_berita = namakomentator.ID_berita
		komentar.Tanggal = namakomentator.Tanggal
		utils.UpdateKomentar(mconn, "komentar", komentar)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil update " + komentar.ID + " dari database"})
	} else {
		c.JSON(http.StatusOK, models.Pesan{Status: false, Message: "Komentar tidak ditemukan"})
	}
}

func HapusKomentar(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var komentar models.Komentar
	err := c.BindJSON(&komentar)
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
	username := c.GetString("username")
	// Cek role
	namakomentator := utils.FindKomentar(mconn, "komentar", komentar)
	if !(role == "admin" || username != namakomentator.Username) {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
		c.Abort()
		return
	}
	// Cek apakah komentar ID ada
	if komentar.ID == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah ID"})
		return
	}
	if utils.IDKomentarExists(mconn, "komentar", komentar) {
		// Delete data komentar
		utils.DeleteKomentar(mconn, "komentar", komentar)
		c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil delete " + komentar.ID + " dari database"})
	} else {
		c.JSON(http.StatusOK, models.Pesan{Status: false, Message: "Komentar tidak ditemukan"})
	}
}
