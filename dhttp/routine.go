package dhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mathyourlife/drips/proto"
)

func (s *HTTPServer) routines() {
	s.mux.HandleFunc("GET /api/routine", s.handleRoutine)
	s.mux.HandleFunc("POST /api/routine", s.handleRoutineCreate)
	s.mux.HandleFunc("DELETE /api/routine/{routineID}", s.handleRoutineDelete)
}

func (s *HTTPServer) Routines() (*proto.RoutinesResponse, error) {
	response, err := s.client.Routines(context.Background(), &proto.RoutinesRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *HTTPServer) handleRoutine(w http.ResponseWriter, r *http.Request) {
	response, err := s.Routines()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send gRPC request: %v", err), http.StatusInternalServerError)
		return
	}

	// Marshal the gRPC response into JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal gRPC response: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func (s *HTTPServer) handleRoutineCreate(w http.ResponseWriter, r *http.Request) {
	// Unmarshal the request body into a protobuf message
	var request proto.RoutineCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	if request.Routine == nil {
		http.Error(w, "Missing routine in request body", http.StatusBadRequest)
		return
	}

	// Send the gRPC request to the server
	response, err := s.client.RoutineCreate(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send gRPC request: %v", err), http.StatusInternalServerError)
		return
	}

	// Marshal the gRPC response into JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal gRPC response: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func (s *HTTPServer) handleRoutineDelete(w http.ResponseWriter, r *http.Request) {
	rIDstr := r.PathValue("routineID")
	rID, err := strconv.Atoi(rIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse routine ID: %v", err), http.StatusBadRequest)
		return
	}

	request := proto.RoutineDeleteRequest{
		RoutineId: int32(rID),
	}
	_, err = s.client.RoutineDelete(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send gRPC request: %v", err), http.StatusInternalServerError)
		return
	}

	response, err := s.Routines()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send gRPC request: %v", err), http.StatusInternalServerError)
		return
	}

	// Marshal the gRPC response into JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal gRPC response: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}
