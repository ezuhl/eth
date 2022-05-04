package main

import (
	"fmt"
	address2 "github.com/eth/internal/address"
	"github.com/eth/internal/data"
	"github.com/eth/internal/handlers"
	routes2 "github.com/eth/internal/routes"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
)

func main() {

	//load env vars
	err := godotenv.Load("../../env/.env")
	if err != nil {
		log.Fatal("could not get environment with error ", err)
	}

	db, err := data.Db()
	if err != nil {
		log.Fatal("could not start db with error ", err)
	}
	//load bus logic
	ethHandler := handlers.NewEthHandler(db)

	//handle routing
	routes := routes2.NewRoutes(ethHandler)

	//find address
	address, err := address2.FindAddress()
	if err != nil {
		log.Fatal("could not get address with error ", err)
	}

	fullAddress := fmt.Sprintf("%s:%d", address, 8080)
	// Run service

	//start listening
	log.Println("start listening on ", fullAddress)
	if err = http.ListenAndServe(fullAddress, routes); err != nil {
		log.Fatal(nil, errors.Wrap(err, "unable to run service"))
	}

	os.Exit(0)
}
