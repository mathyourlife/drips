package dhttp

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/mathyourlife/drips/proto"
)

type HTTPServer struct {
	mux    *http.ServeMux
	conn   *grpc.ClientConn
	client proto.DripsServiceClient
}

func NewHTTPServer(mux *http.ServeMux) (*HTTPServer, error) {
	conn, client, err := grpcClient()
	if err != nil {
		return nil, err
	}

	return &HTTPServer{
		mux:    mux,
		conn:   conn,
		client: client,
	}, nil
}

func (s *HTTPServer) Close() {
	s.conn.Close()
}

func (s *HTTPServer) Start() {
	// Set up object routes
	s.exerciseClassHandlers()
	s.users()

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggingMiddleware(s.mux),
	}
	// Start the HTTP server
	log.Printf("Server listening on %s...", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		d, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(d))

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log after the request is completed
		duration := time.Since(start)
		log.Printf("%s - %s %s - %s - %dms\n",
			r.RemoteAddr,
			r.Method,
			r.URL,
			r.UserAgent(),
			duration.Milliseconds(),
		)
	})
}
