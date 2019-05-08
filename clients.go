package main

import (
	"errors"
	"log"
	"os"

	"cloud.google.com/go/bigtable"
	// "github.com/weathersource/google-cloud-go/bigtable"
	context "golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
)

type Clients struct {
	cbtClient *bigtable.Client
}

func NewClients() (*Clients, error) {

	// get GCP project
	project := "wx-microservices"
	if 0 != len(os.Getenv("is_production")) {
		project = "ws-microservices-production"
	} else if 0 != len(os.Getenv("is_staging")) {
		project = "wx-microservices"
	}

	// get BigTable instance
	instance := "legolas"
	if 0 != len(os.Getenv("bt_ssd_instance")) {
		instance = os.Getenv("bt_ssd_instance")
	}

	// get service account credentials
	var cbtConfig *jwt.Config
	if 0 != len(os.Getenv("password")) {
		var err error
		jsonStr := os.Getenv("password")
		jsonKey := []byte(jsonStr)

		cbtConfig, err = google.JWTConfigFromJSON(jsonKey, bigtable.Scope)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Failed to retrieve service account for BigTable client.")
	}

	// create BigTable client
	log.Println("Configuring cbtClient.")
	cbtClient, err := bigtable.NewClient(
		context.Background(),
		project,
		instance,
		option.WithTokenSource(cbtConfig.TokenSource(context.Background())),
	)
	log.Println("SUCCESS configuring cbtClient.")
	if err != nil {
		return nil, err
	}

	return &Clients{
		cbtClient: cbtClient,
	}, nil
}
