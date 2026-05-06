package api

import (
	"go-notif/internal/api/handler"
	"go-notif/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// rate limit khusus auth
	Limit5 := utils.RateLimiter(10, 1*time.Minute)
	// Limit3 := utils.RateLimiter(3, 1*time.Minute)

	r.GET("/send_wa", Limit5, handler.Send_wa)
	r.GET("/notif", Limit5, handler.Add_notif)

	// r.POST("/oauth", Limit5, handler.Oauth)
	// r.POST("/verify", Limit5, handler.Verify)
	// r.POST("/forgot", Limit3, handler.ForgotPassword)

	// r.GET("/redis", handler.TestRedis)
	// r.GET("/cekredis", handler.DebugSessions)

	// auth := r.Group("/")
	// auth.Use(utils.AuthMiddleware())
	// auth.GET("/decode", handler.Decode)
	// auth.GET("/logout", handler.Logout)
	// auth.GET("/user", Limit5, handler.Get_user_by_id)

	// auth.PUT("/update", controllers.Update_user)
	// auth.POST("/wallet", controllers.Wallet)
}
