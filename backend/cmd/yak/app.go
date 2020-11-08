package yak

import (
	"yak/backend/pkg/handlers"
	"yak/backend/pkg/repositories"
	"yak/backend/pkg/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CreateApp() {
	// logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repositories.NewMongoDB()
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repositories.NewRepository(db)
	services := services.NewService(repos)
	handlers := handlers.NewHandler(services)

	app := fiber.New()
	app.Use(logger.New())
	handlers.RegisterHandlers(app)
	app.Listen(viper.GetString("port"))
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
