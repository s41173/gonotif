// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"context"
	"go-notif/config"
	_ "go-notif/docs"

	"go-notif/internal/api"
	"go-notif/internal/shared"
	"go-notif/internal/worker"

	"os"
	"time"

	"github.com/gin-contrib/cors"

	// "go-notif/internal/auth"
	// chapterService "go-notif/internal/chapter/service"
	// productService "go-notif/internal/product/service"
	// sliderService "go-notif/internal/shared/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load() // load .env
	// set timezone
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	// init logger pertama sebelum yang lain
	config.InitLogger()
	defer config.Log.Sync()

	if err := godotenv.Load(); err != nil {
		config.Log.Warn("no .env file found")
	}

	config.ConnectDatabase()
	config.InitRedis()

	//  set notif property to redis
	var notif shared.NotifProperty
	if err := notif.SetToRedis(config.DB, config.RDB); err != nil {
		config.Log.Fatal("Failed to load notif_property", zap.Error(err))
	}

	// jalankan worker di goroutine terpisah
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker.StartNotifWorker(ctx)

	r := setupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	config.Log.Info("server starting", zap.String("port", port))
	if err := r.Run(":" + port); err != nil {
		config.Log.Fatal("server failed to start", zap.Error(err))
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// CORS - taruh paling atas sebelum route lain
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, // harus false kalau AllowAllOrigins true
		MaxAge:           12 * time.Hour,
	}))

	// handle OPTIONS global
	r.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(204)
	})

	r.Static("/uploads", "./uploads")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "GO NOTIF API SERVICE",
			"status":  "running",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api.RegisterRoutes(r)

	return r
}
