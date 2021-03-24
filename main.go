package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"me2digital.com/event-receiver/routes"
)

//  Add your application configuration to your .env file in the root of your project:
// LISTEN_PORT=...
// ACCESS_LIST=...

// setup two routes, one for shortening the url
// the other for resolving the url
// for example if the short is `4fg`, the user
// must navigate to `localhost:3000/4fg` to redirect to
// original URL. The domain can be changes in .env file
func setupRoutes(app *fiber.App) {
	app.Post("/on_connect", routes.OnConnet)
	app.Post("/on_done", routes.OnDone)
	app.Post("/on_play", routes.OnPlay)
	app.Post("/on_play_done", routes.OnPlayDone)
	app.Post("/on_playlist", routes.OnPlayList)
	app.Post("/on_publish", routes.OnPublish)
	app.Post("/on_publish_done", routes.OnPlayDone)
	app.Post("/on_record_done", routes.OnRecordDone)
	app.Post("/on_record_started", routes.OnRecordStarted)
	app.Post("/on_update", routes.OnUpdate)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	if os.Getenv("MY_CONFIG") == "" {
		panic(fmt.Errorf("No config defined. Please set MY_CONFIG env var!\n"))
	}

	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetConfigName(os.Getenv("MY_CONFIG"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Fatal error ConfigFileNotFoundError: %s \n", err))
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	if viper.GetBool("config.debug") == true {
		fmt.Printf("ENV:MY_CONFIG %#v\n", viper.GetString("MY_CONFIG"))
	}

	if viper.GetBool("config.debug") == true {
		viper.Debug()
	}

	if viper.GetString("config.ListenPort") == "" {
		fmt.Print("I need a ip and port on which I should listen. config.ListenPort\n")
		os.Exit(1)
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:           viper.GetDuration("config.ReadTimeout"),
		WriteTimeout:          viper.GetDuration("config.WriteTimeout"),
		IdleTimeout:           viper.GetDuration("config.IdleTimeout"),
		Concurrency:           viper.GetInt("config.Concurrency"),
		DisableStartupMessage: viper.GetBool("config.DisableStartupMessage"),
	})

	app.Use(logger.New(logger.Config{
		TimeFormat: "02/Jan/2006:15:04:05 -0700",
		Format:     "[${time}] ${status} - ${latency} ${method} ${path} ${bytesSent} ${bytesReceived} ${route} ${error} tcurl:${form:tcurl} ${body}\n",
	}))
	app.Use(recover.New())

	setupRoutes(app)

	if viper.GetBool("config.debug") == true {
		fmt.Printf("Config %#v\n", app.Config())
	}

	log.Fatal(app.Listen(os.Getenv("LISTEN_PORT")))
}
