package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jsagl/go-grpc-from-scratch/api/proto/v1"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartHTTP(ctx context.Context, httpPort string, grpcPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := "localhost:" + grpcPort
	err := v1.RegisterRecipeServiceHandlerFromEndpoint(ctx, mux,  endpoint, opts)
	if err != nil {
		return err
	}

	addr := ":" + httpPort
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
