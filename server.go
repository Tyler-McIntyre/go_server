package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	current_weather "github.com/tyler-mcintyre/go_client/clients/current-weather"
	handlers "github.com/tyler-mcintyre/go_router/api/handlers"
	api "github.com/tyler-mcintyre/go_service_layer/api"
)

func BuildServer() (*http.Server, error) {
	router := chi.NewRouter()
	currentWeatherClient, err := current_weather.NewCurrentWeatherClient()
	if err != nil {
		return nil, err
	}

	service := api.NewService(&currentWeatherClient)

	apiRouter := handlers.NewRouter(router, service)
	apiRouter.AddRoutes()

	port := "8080"
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", "0.0.0.0", port),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server, nil
}
