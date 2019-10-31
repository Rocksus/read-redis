package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rocksus/read-redis/modules/consumer"
	"github.com/Rocksus/read-redis/modules/database"
	"github.com/Rocksus/read-redis/modules/rediscount"
	"github.com/Rocksus/read-redis/modules/server"
	"github.com/Rocksus/read-redis/modules/userdata"
	gorillaHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Init env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//initialize handlers
	redisCountHandler := new(rediscount.Handler)
	redisCountHandler.Init()
	//don't forget to close connection
	defer redisCountHandler.Client.Close()
	userDataHandler := new(userdata.Handler)
	serverHandler := new(server.Handler)
	consumerHandler := new(consumer.Handler)
	consumerHandler.RDC = redisCountHandler
	db := database.ConnectDB()
	//don't forget to close connection
	defer db.Close()
	//pass in variable handlers
	userDataHandler.DB = db
	serverHandler.RDC = redisCountHandler
	serverHandler.UDT = userDataHandler

	//run a consumer goroutine
	go consumerHandler.RunConsumer()
	defer consumerHandler.RDC.Client.Close()

	//Create mux router
	router := mux.NewRouter()
	headers := gorillaHandler.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := gorillaHandler.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := gorillaHandler.AllowedOrigins([]string{"*"})

	router.HandleFunc("/", serverHandler.Serve).Methods("GET")
	fmt.Println("Serving at port 8080")
	log.Fatal(http.ListenAndServe(":8080", gorillaHandler.CORS(headers, methods, origins)(router)))
}
