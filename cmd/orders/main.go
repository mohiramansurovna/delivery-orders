package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mohiramansurovna/orders/internal/application/services"
	"github.com/mohiramansurovna/orders/internal/infra/postgres"
	"github.com/mohiramansurovna/orders/internal/infra/providers"
	"github.com/mohiramansurovna/orders/internal/infra/repositories"
	"github.com/mohiramansurovna/orders/internal/interface/httpapi"
	"github.com/mohiramansurovna/orders/internal/interface/httpapi/controllers"
)

func main() {
	ctx := context.Background()
	pool, err := postgres.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	//infra
	ordersRepository := repositories.NewOrdersRepsitory(pool)
	itemsProvider := providers.NewItemsProvider(os.Getenv("ITEMS_SERVICE_URL"))

	//application
	ordersService := services.NewOrdersService(ordersRepository, itemsProvider)

	//interface
	ordersController := controllers.NewOrdersController(ordersService)

	srv := &http.Server{
		Addr:              os.Getenv("HTTP_ADDR"),
		Handler:           httpapi.NewRouter(pool, ordersController),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("orders service listening on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
