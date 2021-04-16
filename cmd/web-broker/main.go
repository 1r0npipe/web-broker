package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vlslav/web-broker/pkg/rest/endpoint"

	service "github.com/vlslav/web-broker/listening/service"
	file "github.com/vlslav/web-broker/pkg/storage/file"
	_ "github.com/vlslav/web-broker/pkg/storage/memory"
	_ "github.com/vlslav/web-broker/pkg/storage/pgdb"
)

func main() {
	log.Print("Starting the app")

	port := flag.String("port", "8000", "Port")
	storageName := flag.String("storage", "storage.json", "data storage")
	shutdownTimeout := flag.Int64("shutdown_timeout", 3, "shutdown timeout")

	repo := file.NewFileRepo(*storageName)
	svc := service.New(repo)

	serv := http.Server{
		Addr:    net.JoinHostPort("", *port),
		Handler: endpoint.RegisterPublicHTTP(svc),
	}
	go func() {
		if err := serv.ListenAndServe(); err != nil {
			log.Fatalf("listen and serve err: %v", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	log.Print("Started app")

	sig := <-interrupt

	log.Printf("Sig: %v, stopping app", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*shutdownTimeout)*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Printf("shutdown err: %v", err)
	}
}
