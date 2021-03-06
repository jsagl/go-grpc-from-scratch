package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jsagl/go-grpc-from-scratch/api/service/v1"
	"github.com/jsagl/go-grpc-from-scratch/server/grpc"
	"github.com/jsagl/go-grpc-from-scratch/server/rest"
	"github.com/jsagl/go-grpc-from-scratch/storage/postgres"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	//viper.SetConfigFile(".env")
	//if err := viper.ReadInConfig(); err != nil {
	//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	//		fmt.Println("No .env file was found")
	//	} else {
	//		log.Fatalf("Error while reading config file %s", err)
	//	}
	//}


	ctx := context.Background()

	dbAddress := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbAddress)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	recipeStore := storage.NewPostgresRecipeStore(db)
	v1RecipeServiceServer := v1.NewRecipeServiceServer(recipeStore)

	httpPort := os.Getenv("PORT")
	grpcPort := os.Getenv("GRPC_PORT")

	fmt.Println("http" + httpPort)
	fmt.Println("grpc" + grpcPort)

	// run HTTP gateway
	go func() {
		_ = rest.StartHTTP(ctx, httpPort, grpcPort)
	}()

	grpc.StartGRPC(ctx, v1RecipeServiceServer, grpcPort)

}