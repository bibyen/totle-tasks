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
	"github.com/bibyen/totle-tasks/internal/pb/totle_tasks/v1/totletasksv1connect"
)

func main() {
	mux := http.NewServeMux()

	// Register Connect Service
	path, handler := totletasksv1connect.NewGoalServiceHandler(&goalServer{})
	mux.Handle(path, handler)
	path, handler = totletasksv1connect.NewBingoServiceHandler(&bingoServer{})
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
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", port),
		h2c.NewHandler(mux, &http2.Server{}), // Allows us to use HTTP/2 without TLS
	)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// goalServer is minimal implementation for the GoalService
type goalServer struct {
	totletasksv1connect.UnimplementedGoalServiceHandler
}

// bingoServer is a minimal implementation for the BingoService
type bingoServer struct {
	totletasksv1connect.UnimplementedBingoServiceHandler
}
