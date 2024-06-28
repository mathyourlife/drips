package dhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mathyourlife/drips/proto"
)

func (s *HTTPServer) users() {
	s.mux.HandleFunc("GET /api/user", s.handleUser)
	s.mux.HandleFunc("POST /api/user", s.handleUserCreate)
	s.mux.HandleFunc("PUT /api/user/{userID}", s.handleUserUpdate)
	s.mux.HandleFunc("DELETE /api/user/{userID}", s.handleUserDelete)
}

func (s *HTTPServer) Users() (*proto.UsersResponse, error) {
	response, err := s.client.Users(context.Background(), &proto.UsersRequest{})
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *HTTPServer) handleUser(w http.ResponseWriter, r *http.Request) {
	response, err := s.Users()
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

func (s *HTTPServer) handleUserCreate(w http.ResponseWriter, r *http.Request) {
	// Unmarshal the request body into a protobuf message
	var request proto.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	if request.User == nil {
		http.Error(w, "Missing user in request body", http.StatusBadRequest)
		return
	}

	// Send the gRPC request to the server
	response, err := s.client.UserCreate(context.Background(), &request)
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

func (s *HTTPServer) handleUserUpdate(w http.ResponseWriter, r *http.Request) {
	uIDstr := r.PathValue("userID")
	uID, err := strconv.Atoi(uIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse user ID: %v", err), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body into a protobuf message
	var request proto.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	if request.User == nil {
		http.Error(w, "Missing user in request body", http.StatusBadRequest)
		return
	}
	request.User.UserId = int32(uID)

	// Send the gRPC request to the server
	response, err := s.client.UserUpdate(context.Background(), &request)
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

func (s *HTTPServer) handleUserDelete(w http.ResponseWriter, r *http.Request) {
	uIDstr := r.PathValue("userID")
	uID, err := strconv.Atoi(uIDstr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse user ID: %v", err), http.StatusBadRequest)
		return
	}

	request := proto.UserDeleteRequest{
		UserId: int32(uID),
	}
	_, err = s.client.UserDelete(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send gRPC request: %v", err), http.StatusInternalServerError)
		return
	}

	response, err := s.Users()
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
