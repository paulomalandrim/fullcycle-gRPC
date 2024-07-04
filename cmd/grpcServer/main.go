package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/paulomalandrim/fullcycle-gRPC/internal/database"
	"github.com/paulomalandrim/fullcycle-gRPC/internal/pb"
	"github.com/paulomalandrim/fullcycle-gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	log.Println("Starting server...")

	log.Println("Connecting to database...")
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	log.Println("Starting gRPC server...")
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	port := 50051
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		panic(err)
	}

	log.Printf("Server started on port %d", port)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
