package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/mathyourlife/drips/proto"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	dev := os.Getenv("DEV") == "true"
	migrationSourceURL := os.Getenv("MIGRATION_SOURCE_URL")

	err := initDB(dbPath, migrationSourceURL)
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer dbHandle.Close()

	// Create a gRPC server
	grpcServer := grpc.NewServer()
	proto.RegisterDripsServiceServer(grpcServer, &DripsServer{})

	// Start the gRPC server on a port
	go func() {
		listen, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("net listen failed: %s", err)
		}
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Create an HTTP server
	mux := http.NewServeMux()

	if dev {
		log.Println("Running in development mode, setting up proxy to locally running dev React app")
		targetURL, err := url.Parse("http://localhost:3000")
		if err != nil {
			log.Fatal(err)
		}
		// Create a reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		// Set up the reverse proxy for all other requests
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Println("Proxying request to React app")
			proxy.ServeHTTP(w, r)
		})
	} else {
		mux.Handle("/", http.FileServer(http.Dir("./frontend/build"))) // Serve the built React app
	}

	httpServer, err := NewHTTPServer(mux)
	if err != nil {
		log.Fatalf("failed to create HTTP server: %v", err)
	}
	defer httpServer.Close()

	httpServer.Start()
}

func grpcClient() (*grpc.ClientConn, proto.DripsServiceClient, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to gRPC server: %s", err)
	}
	// Create a gRPC client
	client := proto.NewDripsServiceClient(conn)
	return conn, client, nil
}
