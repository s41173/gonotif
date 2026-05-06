package handler

import (
	"go-notif/internal/api/service"

	"github.com/gin-gonic/gin"
)

func Send_wa(c *gin.Context) {
	stts := service.Send_wa("082277014410", "Hello From Go")
	if !stts {
		c.JSON(400, gin.H{"error": "Gagal Kirim Pesan"})
		return
	}
	c.JSON(200, gin.H{"message": "Pesan Berhasil Dikirim"})
}

func Add_notif(c *gin.Context) {
	stts, err := service.Add_notif(422, "Jay Kiran", "082277014410", "Test Notif", "Isi Content", 7, "login", 0)
	if err != nil {
		c.JSON(stts, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Notif Berhasil Ditambahkan"})
}
