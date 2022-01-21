package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Version     string
	Environment string
}

type Application struct {
	Config Config
	Logger *log.Logger
}

func main() {

	//Read .env file
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error :", err)
	}

	//Config setting
	var config Config

	flag.StringVar(&config.Port, "port", envs["PORT"], "Server port")
	flag.StringVar(&config.Version, "version", "1.0", "Program version")
	flag.StringVar(&config.Environment, "environment", "Development", "Server environment")
	flag.Parse()

	//setting application
	app := Application{
		Config: config,
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	//Connecting to route and serve
	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", config.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Println("Server is running on port:", config.Port)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal("Error:", err)
	}

}
