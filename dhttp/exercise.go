package dhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mathyourlife/drips/proto"
)

func (s *HTTPServer) exercises() {
	s.mux.HandleFunc("GET /api/exercise", s.handleExercise)
	s.mux.HandleFunc("POST /api/exercise", s.handleExerciseCreate)
	s.mux.HandleFunc("DELETE /api/exercise/{exerciseID}", s.handleExerciseDelete)
}

func (s *HTTPServer) Exercises() (*proto.ExercisesResponse, error) {
	response, err := s.client.Exercises(context.Background(), &proto.ExercisesRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *HTTPServer) handleExercise(w http.ResponseWriter, r *http.Request) {
	response, err := s.Exercises()
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

func (s *HTTPServer) handleExerciseCreate(w http.ResponseWriter, r *http.Request) {
	// Unmarshal the request body into a protobuf message
	var request proto.ExerciseCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	if request.Exercise == nil {
		http.Error(w, "Missing exercise in request body", http.StatusBadRequest)
		return
	}

	// Send the gRPC request to the server
	response, err := s.client.ExerciseCreate(context.Background(), &request)
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

func (s *HTTPServer) handleExerciseDelete(w http.ResponseWriter, r *http.Request) {
	eIDstr := r.PathValue("exerciseID")
	eID, err := strconv.Atoi(eIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse exercise ID: %v", err), http.StatusBadRequest)
		return
	}

	request := proto.ExerciseDeleteRequest{
		ExerciseId: int32(eID),
	}
	_, err = s.client.ExerciseDelete(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send gRPC request: %v", err), http.StatusInternalServerError)
		return
	}

	response, err := s.Exercises()
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
