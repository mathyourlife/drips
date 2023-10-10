package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mathyourlife/drips"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/mathyourlife/drips/model"
	pb "github.com/mathyourlife/drips/proto"
)

func main() {
	fmt.Println("drips server")

	app := cli.NewApp()
	app.Name = "drips-server"
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start a drips server",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "db-file", Required: true},
				cli.StringFlag{Name: "addr", Value: "localhost:5050"},
			},
			Action: func(c *cli.Context) error {
				return startServer(c)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func startServer(c *cli.Context) error {
	db, err := gorm.Open(sqlite.Open(c.String("db-file")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect database")
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

	log.Printf("starting server on %s", c.String("addr"))
	lis, err := net.Listen("tcp", c.String("addr"))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	return server.Serve(lis)
}
