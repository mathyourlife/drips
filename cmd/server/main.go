package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mathyourlife/drips"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/mathyourlife/drips/model"
	pb "github.com/mathyourlife/drips/proto"
)

func main() {
	fmt.Println("starting server")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	defer func() {
		d, _ := db.DB()
		d.Close()
	}()

	db.Debug().AutoMigrate(
		&model.User{},
		&model.Modifier{},
		&model.ExerciseClass{},
		&model.Exercise{},
		&model.Routine{},
	)

	svc := drips.NewService(db)
	server := grpc.NewServer()
	pb.RegisterDripsServiceServer(server, svc)

	lis, err := net.Listen("tcp", "localhost:5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server.Serve(lis)
}
