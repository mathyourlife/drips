package dgrpc

import (
	"context"

	"github.com/mathyourlife/drips/proto"
)

// Implement your gRPC methods
func (s *DripsServer) ModifierCreate(ctx context.Context, req *proto.ModifierCreateRequest) (*proto.ModifierCreateResponse, error) {
	// Create the modifier in the database

	result, err := s.db.Exec(`
        INSERT INTO modifier (name)
        VALUES (?)
        `, req.Modifier.Name)
	if err != nil {
		return nil, err
	}

	// parse the result
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Retrieve the inserted record
	var m proto.Modifier
	err = s.db.QueryRow(`
        SELECT
                modifier_id, name
        FROM modifier
        WHERE modifier_id = ?`, lastInsertID).Scan(
		&m.ModifierId,
		&m.Name,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ModifierCreateResponse{
		Modifier: &m,
	}, nil
}

// Implement your gRPC methods
func (s *DripsServer) Modifiers(ctx context.Context, req *proto.ModifiersRequest) (*proto.ModifiersResponse, error) {
	// List all modifiers

	// Return the list of modifiers
	var ms []*proto.Modifier
	rows, err := s.db.Query(`
        SELECT
                modifier_id, name
        FROM modifier`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var m proto.Modifier
		err := rows.Scan(&m.ModifierId, &m.Name)
		if err != nil {
			return nil, err
		}
		ms = append(ms, &m)
	}

	return &proto.ModifiersResponse{
		Modifiers: ms,
	}, nil
}

// Delete the modifier from the database
func (s *DripsServer) ModifierDelete(ctx context.Context, req *proto.ModifierDeleteRequest) (*proto.ModifierDeleteResponse, error) {
	_, err := s.db.Exec(`
        DELETE FROM modifier
        WHERE modifier_id = ?
        `, req.ModifierId)
	if err != nil {
		return nil, err
	}

	return &proto.ModifierDeleteResponse{}, nil
}
