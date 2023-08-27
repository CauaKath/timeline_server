package main

import (
	"fmt"
	"log"

	"github.com/cauakath/timeline-server/config"
	"github.com/cauakath/timeline-server/controller"
	"github.com/cauakath/timeline-server/database"
	"github.com/cauakath/timeline-server/model"
	"github.com/cauakath/timeline-server/repo"
	"github.com/cauakath/timeline-server/router"
	"github.com/cauakath/timeline-server/usecase"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Hello, World!")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load enviroment variables", err)
	}

	db := database.ConnectionDB(&loadConfig)

	db.AutoMigrate(&model.Timeline{})

	rdb := database.ConnectionRedisDb(&loadConfig)

	startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()

	timelineRepo := repo.NewTimelineRepo(db, rdb)
	timelineUseCase := usecase.NewTimelineUseCase(timelineRepo)
	timelineController := controller.NewTimelineController(timelineUseCase)

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Access-Control-Allow-Origin",
	}))

	routes := router.NewRouter(app, timelineController)

	err := routes.Listen(":3400")
	if err != nil {
		panic(err)
	}
}
