package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mohiramansurovna/orders/internal/infra/postgres"
	"github.com/mohiramansurovna/orders/internal/interface/httpapi"
	
)

func main() {
	ctx:=context.Background()
	pool, err :=postgres.Connect(ctx,os.Getenv("DATABASE_URL"))
	if err!=nil{
		log.Fatal(err)
	}
	defer pool.Close()

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           httpapi.NewRouter(pool),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("orders service listening on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
