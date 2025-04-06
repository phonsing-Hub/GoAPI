package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	"github.com/phonsing-Hub/GoAPI/config"
	_ "github.com/phonsing-Hub/GoAPI/docs"
	"github.com/phonsing-Hub/GoAPI/entity"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/subosito/gotenv"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


func init() {
	_ = gotenv.Load()
}

// @title 						Go Skeleton!
// @version 					1.0
// @description 				This is a sample swagger for Go Skeleton
// @termsOfService 				http://swagger.io/terms/
// @contact.name 				API Support
// @contact.email 				rahmat.putra@spesolution.com
// @license.name				Apache 2.0
// @securityDefinitions.apikey 	Bearer
// @in							header
// @name						Authorization
// @license.url 				http://www.apache.org/licenses/LICENSE-2.0.html
// @host 						localhost:7011
// @BasePath /
func main() {
	// Initialize config variable from .env file
	cfg := config.NewConfig()

	app := fiber.New(config.NewFiberConfiguration(cfg))
	app.Get("/apidoc/*", swagger.HandlerDefault)

	setupMiddleware(app, cfg)



	gormLogger := config.NewGormLogPostgreConfig(&cfg.PostgreSqlOption)
	postgreDB, err := config.NewPostgreSQL(cfg.AppEnv, &cfg.PostgreSqlOption, gormLogger)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PostgreSQL connected: ", postgreDB)
	postgreDB.DB.AutoMigrate(&User{})


	app.Get("/health-check", healthCheck)
	app.Get("/metrics", monitor.New())

	// Handle Route not found
	app.Use(routeNotFound)

	runServerWithGracefulShutdown(app, cfg.ApiPort, 30)
}

func setupMiddleware(app *fiber.App, cfg *config.Config) {
	// Enable CORS if API shared in public
	// if cfg.AppEnv == "production" {
	// 	app.Use(
	// 		cors.New(cors.Config{
	// 			AllowCredentials: true,
	// 			AllowOrigins:     cfg.AllowedCredentialOrigins,
	// 			AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	// 			AllowMethods:     "GET,POST,PUT,DELETE,PATCH",
	// 		}),
	// 	)
	// }

	app.Use(
		logger.New(logger.Config{
			Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
			TimeFormat: "04-Apr-2002 00:00:00",
			TimeZone:   "Asia/Bangkok",
		}),
		recover.New(recover.Config{
			StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
				fmt.Println(c.Request().URI())
				stacks := fmt.Sprintf("panic: %v\n%s\n", e, debug.Stack())
				log.Println(stacks)
			},
			EnableStackTrace: true,
		}),
	)
}

func runServerWithGracefulShutdown(app *fiber.App, apiPort string, shutdownTimeout int) {
	var wg sync.WaitGroup
	wg.Add(1)

	// Run server in a goroutine
	go func() {
		defer wg.Done()
		log.Printf("Starting REST server, listening at %s\n", apiPort)
		if err := app.Listen(apiPort); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Capture OS signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down REST server...")

	// Timeout context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(shutdownTimeout)*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	} else {
		log.Println("REST server shut down gracefully")
	}

	// Wait for goroutines to exit
	wg.Wait()
	log.Println("All tasks completed. Exiting application.")
}

var healthCheck = func(c *fiber.Ctx) error {
	return c.JSON(entity.GeneralResponse{
		Code:    200,
		Message: "OK!",
	})
}

var routeNotFound = func(c *fiber.Ctx) error {
	return c.Status(404).JSON(entity.GeneralResponse{
		Code:    404,
		Message: "Route Not Found!",
	})
}
