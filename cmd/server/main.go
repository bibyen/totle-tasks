package main

import (
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// Replace with your actual module path
	"github.com/bibyen/totle-tasks/internal/domain"
	"github.com/bibyen/totle-tasks/internal/pb/totle_tasks/v1/totletasksv1connect"
	"github.com/bibyen/totle-tasks/internal/server"
)

func main() {

	// Initialise server
	goalService := domain.GoalService{}
	bingoService := domain.BingoService{}
	server, err := server.NewServer(&goalService, &bingoService)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	mux := http.NewServeMux()

	// Register Connect Service
	path, handler := totletasksv1connect.NewGoalServiceHandler(server)
	mux.Handle(path, handler)
	path, handler = totletasksv1connect.NewBingoServiceHandler(server)
	mux.Handle(path, handler)

	// Register Health Service - to report service health status
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(
			"totle_tasks.v1.GoalService",
			"totle_tasks.v1.BingoService",
		),
	))

	// Register reflection - to allow tools like Postman or 'buf curl' to explore the API
	reflector := grpcreflect.NewStaticReflector(
		"totle_tasks.v1.GoalService",
		"totle_tasks.v1.BingoService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// Start the H2C Server
	port := "8080"
	log.Printf("Starting Connect server on :%s", port)
	err = http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		h2c.NewHandler(mux, &http2.Server{}), // Allows us to use HTTP/2 without TLS
	)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
