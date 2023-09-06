package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func router(mux *http.ServeMux) {
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/webhook", hook)
}

func main() {
	mux := http.NewServeMux()
	router(mux)
	srv := &http.Server{
		Addr:              ":" + Config.Port,
		Handler:           mux,
		ReadHeaderTimeout: 1 * time.Second,
		MaxHeaderBytes:    1 << 17, // 128 KB
	}

	go func() {
		log.Printf("start server, port: %v", Config.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	stopCh := make(chan os.Signal, 1)
	defer close(stopCh)

	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	log.Println("shutdown server with Signal ", <-stopCh)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
	log.Println("server shutdown")
}
