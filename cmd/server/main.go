package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/mathyourlife/drips"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"

	pb "github.com/mathyourlife/drips/proto"
)

func main() {
	fmt.Println("starting server")

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmts := []string{
		`CREATE TABLE user (user_id BIGSERIAL PRIMARY KEY, display_name TEXT);`,
		`CREATE TABLE routine (routine_id BIGSERIAL PRIMARY KEY, name TEXT, source TEXT, sequence INT);`,
		`CREATE TABLE class (class_id BIGSERIAL PRIMARY KEY, name TEXT, short_name TEXT);`,
		`CREATE TABLE exercise (exercise_id BIGSERIAL PRIMARY KEY, sequence INT, class_id INT, duration_seconds INT, rest_seconds INT);`,
		`CREATE TABLE routine_exercise (routine_exercise_id BIGSERIAL PRIMARY KEY, routine_id INT, exercise_id INT);`,
		`CREATE TABLE modifier (modifier_id BIGSERIAL PRIMARY KEY, name TEXT);`,
		`CREATE TABLE exercise_modifier (exercise_modifier_id BIGSERIAL PRIMARY KEY, exercise_id INT, modifier_id INT);`,
		`INSERT INTO routine (routine_id, name, source, sequence) VALUES (4, 'my workout', 'https://localhost', 4), (5, 'neighborhood run', 'https://localhost', 0);`,
	}

	for _, stmt := range stmts {
		_, err = db.Exec(stmt)
		if err != nil {
			log.Printf("%q: %s", err, stmt)
			return
		}
	}

	svc := drips.NewService(db)
	server := grpc.NewServer()
	pb.RegisterDripsServiceServer(server, svc)

	lis, err := net.Listen("tcp", "localhost:5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server.Serve(lis)

}
