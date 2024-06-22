package main

import (
	"context"

	"github.com/mathyourlife/drips/proto"
)

// Implement your gRPC methods
func (s *DripsServer) UserCreate(ctx context.Context, req *proto.UserCreateRequest) (*proto.UserCreateResponse, error) {
	// Create the user in the database

	result, err := dbHandle.Exec(`
	INSERT INTO user (first_name, last_name)
	VALUES (?, ?)
	`, req.User.FirstName, req.User.LastName)
	if err != nil {
		return nil, err
	}

	// parse the result
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Retrieve the inserted record
	var u proto.User
	err = dbHandle.QueryRow(`
	SELECT
		user_id, first_name, last_name
	FROM user
	WHERE user_id = ?`, lastInsertID).Scan(
		&u.UserId,
		&u.FirstName,
		&u.LastName,
	)
	if err != nil {
		return nil, err
	}

	return &proto.UserCreateResponse{
		User: &u,
	}, nil
}

// Implement your gRPC methods
func (s *DripsServer) Users(ctx context.Context, req *proto.UsersRequest) (*proto.UsersResponse, error) {
	// List all users

	// Return the list of users
	var us []*proto.User
	rows, err := dbHandle.Query(`
	SELECT
		user_id, first_name, last_name
	FROM user`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u proto.User
		err := rows.Scan(&u.UserId, &u.FirstName, &u.LastName)
		if err != nil {
			return nil, err
		}
		us = append(us, &u)
	}

	return &proto.UsersResponse{
		Users: us,
	}, nil
}

// Delete the user from the database
func (s *DripsServer) UserDelete(ctx context.Context, req *proto.UserDeleteRequest) (*proto.UserDeleteResponse, error) {
	_, err := dbHandle.Exec(`
	DELETE FROM user
	WHERE user_id = ?
	`, req.UserId)
	if err != nil {
		return nil, err
	}

	return &proto.UserDeleteResponse{}, nil
}
