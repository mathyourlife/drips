package dhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mathyourlife/drips/proto"
)

func (s *HTTPServer) exerciseClassHandlers() {
	s.mux.HandleFunc("GET /api/exercise_class", s.handleExerciseClass)
	s.mux.HandleFunc("POST /api/exercise_class", s.handleExerciseClassCreate)
	s.mux.HandleFunc("DELETE /api/exercise_class/{exerciseClassID}", s.handleExerciseClassDelete)
}

func (s *HTTPServer) ExerciseClasses() (*proto.ExerciseClassesResponse, error) {
	response, err := s.client.ExerciseClasses(context.Background(), &proto.ExerciseClassesRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *HTTPServer) handleExerciseClass(w http.ResponseWriter, r *http.Request) {
	response, err := s.ExerciseClasses()
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

func (s *HTTPServer) handleExerciseClassCreate(w http.ResponseWriter, r *http.Request) {
	// Unmarshal the request body into a protobuf message
	var request proto.ExerciseClassCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	if request.ExerciseClass == nil {
		http.Error(w, "Missing exercise class in request body", http.StatusBadRequest)
		return
	}

	// Send the gRPC request to the server
	response, err := s.client.ExerciseClassCreate(context.Background(), &request)
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

func (s *HTTPServer) handleExerciseClassDelete(w http.ResponseWriter, r *http.Request) {
	ecIDstr := r.PathValue("exerciseClassID")
	ecID, err := strconv.Atoi(ecIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse exercise class ID: %v", err), http.StatusBadRequest)
		return
	}

	request := proto.ExerciseClassDeleteRequest{
		ExerciseClassId: int32(ecID),
	}
	_, err = s.client.ExerciseClassDelete(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send gRPC request: %v", err), http.StatusInternalServerError)
		return
	}

	response, err := s.ExerciseClasses()
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
