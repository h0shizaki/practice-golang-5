package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/models"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Port        string
	Version     string
	Environment string
	db          struct {
		dsn string
	}
}

type Application struct {
	Config Config
	Logger *log.Logger
	Models models.Models
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
	flag.StringVar(&config.db.dsn, "database", envs["DB"], "Server database")
	flag.Parse()

	//setting application
	app := Application{
		Config: config,
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}

	//Connecting to database
	db, err := openDB(config)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer db.Close()

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

// Function that create a connection to database

func openDB(cfg Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.db.dsn)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	fmt.Println("DB connected")
	return db, nil

}
