package dhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mathyourlife/drips/proto"
)

func (s *HTTPServer) exerciseModifiers() {
	s.mux.HandleFunc("GET /api/exercise_modifier", s.handleExerciseModifier)
	s.mux.HandleFunc("POST /api/exercise_modifier", s.handleExerciseModifierCreate)
	s.mux.HandleFunc("GET /api/exercise_modifier/{exerciseModifierID}", s.handleExerciseModifierGet)
	s.mux.HandleFunc("PUT /api/exercise_modifier/{exerciseModifierID}", s.handleExerciseModifierUpdate)
	s.mux.HandleFunc("DELETE /api/exercise_modifier/{exerciseModifierID}", s.handleExerciseModifierDelete)
}

func (s *HTTPServer) ExerciseModifiers() (*proto.ExerciseModifiersResponse, error) {
	response, err := s.client.ExerciseModifiers(context.Background(), &proto.ExerciseModifiersRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *HTTPServer) handleExerciseModifierGet(w http.ResponseWriter, r *http.Request) {
	emIDstr := r.PathValue("exerciseModifierID")
	emID, err := strconv.Atoi(emIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse exercise modifier ID: %v", err), http.StatusBadRequest)
		return
	}

	response, err := s.client.ExerciseModifier(context.Background(), &proto.ExerciseModifierRequest{
		ExerciseModifierId: int32(emID),
	})
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

func (s *HTTPServer) handleExerciseModifier(w http.ResponseWriter, r *http.Request) {
	response, err := s.ExerciseModifiers()
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

func (s *HTTPServer) handleExerciseModifierCreate(w http.ResponseWriter, r *http.Request) {
	// Unmarshal the request body into a protobuf message
	var request proto.ExerciseModifierCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	if request.ExerciseModifier == nil {
		http.Error(w, "Missing exercise modifier in request body", http.StatusBadRequest)
		return
	}

	// Send the gRPC request to the server
	response, err := s.client.ExerciseModifierCreate(context.Background(), &request)
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

func (s *HTTPServer) handleExerciseModifierUpdate(w http.ResponseWriter, r *http.Request) {
	emIDstr := r.PathValue("exerciseModifierID")
	emID, err := strconv.Atoi(emIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse exercise modifier ID: %v", err), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body into a protobuf message
	var request proto.ExerciseModifierUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	if request.ExerciseModifier == nil {
		http.Error(w, "Missing exercise modifier in request body", http.StatusBadRequest)
		return
	}
	request.ExerciseModifier.ExerciseModifierId = int32(emID)

	// Send the gRPC request to the server
	response, err := s.client.ExerciseModifierUpdate(context.Background(), &request)
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

func (s *HTTPServer) handleExerciseModifierDelete(w http.ResponseWriter, r *http.Request) {
	emIDstr := r.PathValue("exerciseModifierID")
	emID, err := strconv.Atoi(emIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse exercise modifier ID: %v", err), http.StatusBadRequest)
		return
	}

	request := proto.ExerciseModifierDeleteRequest{
		ExerciseModifierId: int32(emID),
	}
	_, err = s.client.ExerciseModifierDelete(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send gRPC request: %v", err), http.StatusInternalServerError)
		return
	}

	response, err := s.ExerciseModifiers()
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
