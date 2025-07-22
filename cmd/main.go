package main

import (
	"flag"
	"fmt"
	"log"

	_ "test-assignment/docs"
	"test-assignment/internal/app/db"
	"test-assignment/internal/app/handler"
	"test-assignment/internal/app/repository"
	"test-assignment/internal/app/server"
	"test-assignment/internal/app/services"

	"github.com/joho/godotenv"
)

func main() {

	migr := flag.Bool("migrate", false, fmt.Sprint("Migrating Entity"))
	start := flag.Bool("start", false, fmt.Sprint("Starting server"))
	docker := flag.Bool("docker", false, fmt.Sprint("Loading docker env"))
	flag.Parse()
	if *docker {
		err := godotenv.Load("docker.env")
		if err != nil {
			log.Fatal("Error loading docker.env file")
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	if *migr {
		session := db.Connect()
		log.Println("Migrating")
		db.Migrate(session)
	}
	if *start {
		srv := new(server.Server)
		log.Println("Starting server")
		defer log.Println("End of Program")
		log.Println("Open the server")
		fmt.Println("Running and Serving on: http://127.0.0.1:5000")
		session := db.Connect()
		repo := repository.SubscriptionRepo{DB: session}
		svc := services.SubscriptionService{Repo: repo}
		mHandler := handler.SubscriptionHandler{Svc: svc}
		if err := srv.Run(":5000", mHandler.InitRoutes()); err != nil {
			log.Fatalln(err)
		}
	}
}
