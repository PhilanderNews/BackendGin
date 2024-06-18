package controllers

import (
	"net/http"
	"os"

	"github.com/PhilanderNews/BackendGin/helpers"
	"github.com/PhilanderNews/BackendGin/middleware"
	"github.com/PhilanderNews/BackendGin/models"
	"github.com/PhilanderNews/BackendGin/utils"
	"github.com/gin-gonic/gin"
)

func Registrasi(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var user models.Users
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	// Cek apakah username telah dipakai
	if utils.UsernameExists(mconn, "users", user) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Username telah dipakai"})
		return
	}
	// Cek apakah kode verifikasi benar
	if user.Kode == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Masukkan kode verifikasi"})
		return
	}
	dataverifikasi := utils.FindVerifikasi(mconn, "verifikasi", user)
	if user.Kode != dataverifikasi.Kode {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Kode tidak sama"})
		return
	}
	// Hash password
	hash, hashErr := helpers.HashPassword(user.Password)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Gagal hash password: " + hashErr.Error()})
		return
	}
	user.Password = hash
	// Insert data user
	utils.InsertUser(mconn, "users", user)
	c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil input data"})
}

func Login(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var user models.Users
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	// Cek apakah username ada
	if !utils.UsernameExists(mconn, "users", user) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Akun tidak ditemukan"})
		return
	}
	// Cek password
	if !utils.IsPasswordValid(mconn, "users", user) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Password salah"})
		return
	}
	// Encode data login
	datauser := utils.FindUser(mconn, "users", user)
	tokenstring, tokenerr := helpers.Encode(datauser.Name, datauser.Username, datauser.Role, os.Getenv("privatekey"))
	if tokenerr != nil {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Gagal encode token: " + tokenerr.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil login", Token: tokenstring})
}

func AmbilSatuUser(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var user models.Users
	err := c.BindJSON(&user)
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
	// Cek role
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
		c.Abort()
		return
	}
	// Get data user
	datauser := utils.FindUser(mconn, "users", user)
	c.JSON(http.StatusOK, datauser)
}

func AmbilSemuaUser(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	// Authorization
	middleware.Authorization(c)
	if c.IsAborted() {
		return
	}
	role := c.GetString("role")
	// Cek role
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
		c.Abort()
		return
	}
	// Get data user
	datauser, err := utils.GetAllUser(mconn, "users")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "GetAllDoc error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, datauser)
}

func EditUser(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var user models.Users
	err := c.BindJSON(&user)
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
	// Cek role
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
		c.Abort()
		return
	}
	// Cek apakah username ada
	if user.Username == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah username"})
		return
	}
	if !utils.UsernameExists(mconn, "users", user) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Akun yang ingin diedit tidak ditemukan"})
		return
	}
	// Hash password jika password di isi
	if user.Password != "" {
		hash, hashErr := helpers.HashPassword(user.Password)
		if hashErr != nil {
			c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Gagal hash password: " + hashErr.Error()})
			return
		}
		user.Password = hash
	} else {
		datauser := utils.FindUser(mconn, "users", user)
		user.Password = datauser.Password
	}
	// Update data user
	utils.UpdateUser(mconn, "users", user)
	c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil update " + user.Username + " dari database"})
}

func HapusUser(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var user models.Users
	err := c.BindJSON(&user)
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
	// Cek role
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, models.Pesan{Status: false, Message: "Anda tidak memiliki akses"})
		c.Abort()
		return
	}
	// Cek apakah username ada
	if user.Username == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Parameter dari function ini adalah username"})
		return
	}
	if !utils.UsernameExists(mconn, "users", user) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Akun yang ingin dihapus tidak ditemukan"})
		return
	}
	// Delete data user
	utils.DeleteUser(mconn, "users", user)
	c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berhasil hapus " + user.Username + " dari database"})
}
