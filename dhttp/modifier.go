package dhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mathyourlife/drips/proto"
)

func (s *HTTPServer) modifiers() {
	s.mux.HandleFunc("GET /api/modifier", s.handleModifier)
	s.mux.HandleFunc("POST /api/modifier", s.handleModifierCreate)
	s.mux.HandleFunc("PUT /api/modifier/{modifierID}", s.handleModifierUpdate)
	s.mux.HandleFunc("DELETE /api/modifier/{modifierID}", s.handleModifierDelete)
}

func (s *HTTPServer) Modifiers() (*proto.ModifiersResponse, error) {
	response, err := s.client.Modifiers(context.Background(), &proto.ModifiersRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *HTTPServer) handleModifier(w http.ResponseWriter, r *http.Request) {
	response, err := s.Modifiers()
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

func (s *HTTPServer) handleModifierCreate(w http.ResponseWriter, r *http.Request) {
	// Unmarshal the request body into a protobuf message
	var request proto.ModifierCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	if request.Modifier == nil {
		http.Error(w, "Missing modifier in request body", http.StatusBadRequest)
		return
	}

	// Send the gRPC request to the server
	response, err := s.client.ModifierCreate(context.Background(), &request)
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

func (s *HTTPServer) handleModifierUpdate(w http.ResponseWriter, r *http.Request) {
	modIDstr := r.PathValue("modifierID")
	modID, err := strconv.Atoi(modIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse modifier ID: %v", err), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body into a protobuf message
	var request proto.ModifierUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	if request.Modifier == nil {
		http.Error(w, "Missing modifier in request body", http.StatusBadRequest)
		return
	}
	request.Modifier.ModifierId = int32(modID)

	// Send the gRPC request to the server
	response, err := s.client.ModifierUpdate(context.Background(), &request)
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

func (s *HTTPServer) handleModifierDelete(w http.ResponseWriter, r *http.Request) {
	mIDstr := r.PathValue("modifierID")
	mID, err := strconv.Atoi(mIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse modifier ID: %v", err), http.StatusBadRequest)
		return
	}

	request := proto.ModifierDeleteRequest{
		ModifierId: int32(mID),
	}
	_, err = s.client.ModifierDelete(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send gRPC request: %v", err), http.StatusInternalServerError)
		return
	}

	response, err := s.Modifiers()
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
