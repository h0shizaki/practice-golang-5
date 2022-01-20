package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Version     string
	Environment string
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	fmt.Println("Server is running on port:", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil)

}
