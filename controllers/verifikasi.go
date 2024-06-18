package controllers

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/PhilanderNews/BackendGin/models"
	"github.com/PhilanderNews/BackendGin/utils"
	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/gin-gonic/gin"
	"github.com/whatsauth/wa"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) (string, error) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}
	return string(result), nil
}

func GenerateVerifikasi(c *gin.Context) {
	mconn := utils.SetConnection("mongoenv", "philandernews")
	var verifikasi models.Verifikasi
	var user models.Users
	err := c.BindJSON(&verifikasi)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Pesan{Status: false, Message: "Error parsing application/json: " + err.Error()})
		return
	}
	// Cek apakah username ada
	if utils.UsernameExists(mconn, "users", user) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Username telah dipakai"})
		return
	}
	// Cek apakah nomor hp diisi
	if verifikasi.NoWa == "" {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Isi nomor whatsapp anda"})
		return
	}
	// Cek apakah kode telah dibuat
	if utils.VerifikasiExists(mconn, "verifikasi", verifikasi) {
		c.JSON(http.StatusInternalServerError, models.Pesan{Status: false, Message: "Kode telah dikirim, tunggu 5 menit untuk generate kode lagi"})
		return
	}
	str, _ := randomString(12)
	verifikasi.Kode = str
	verifikasi.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	// Insert data user
	utils.InsertVerifikasi(mconn, "verifikasi", verifikasi)
	// Prepare and send a WhatsApp message with registration details
	var nohp = user.NoWa

	dt := &wa.TextMessage{
		To:       nohp,
		IsGroup:  false,
		Messages: "Berikut kode verifikasi anda: " + str,
	}

	// Make an API call to send WhatsApp message
	atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("tokenwa"), dt, "https://api.wa.my.id/api/send/message/text")
	c.JSON(http.StatusOK, models.Pesan{Status: true, Message: "Berikut kode verifikasi anda: " + verifikasi.Kode})
}
