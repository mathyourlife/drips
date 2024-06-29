package dgrpc

import (
	"context"

	"github.com/mathyourlife/drips/proto"
)

// Implement your gRPC methods
func (s *DripsServer) RoutineCreate(ctx context.Context, req *proto.RoutineCreateRequest) (*proto.RoutineCreateResponse, error) {
	// Create the routine in the database

	result, err := s.db.Exec(`
        INSERT INTO routine (name, source)
        VALUES (?, ?)
        `, req.Routine.Name, req.Routine.Source)
	if err != nil {
		return nil, err
	}

	// parse the result
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Retrieve the inserted record
	var r proto.Routine
	err = s.db.QueryRow(`
        SELECT
			routine_id, name, source
        FROM routine
        WHERE routine_id = ?`, lastInsertID).Scan(
		&r.RoutineId,
		&r.Name,
		&r.Source,
	)
	if err != nil {
		return nil, err
	}

	return &proto.RoutineCreateResponse{
		Routine: &r,
	}, nil
}

// Implement your gRPC methods
func (s *DripsServer) Routines(ctx context.Context, req *proto.RoutinesRequest) (*proto.RoutinesResponse, error) {
	// List all routines

	// Return the list of routines
	var rs []*proto.Routine
	rows, err := s.db.Query(`
        SELECT
			routine_id, name, source
        FROM routine`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var r proto.Routine
		err := rows.Scan(&r.RoutineId, &r.Name, &r.Source)
		if err != nil {
			return nil, err
		}
		rs = append(rs, &r)
	}

	return &proto.RoutinesResponse{
		Routines: rs,
	}, nil
}

// Delete the routine from the database
func (s *DripsServer) RoutineDelete(ctx context.Context, req *proto.RoutineDeleteRequest) (*proto.RoutineDeleteResponse, error) {
	_, err := s.db.Exec(`
        DELETE FROM routine
        WHERE routine_id = ?
        `, req.RoutineId)
	if err != nil {
		return nil, err
	}

	return &proto.RoutineDeleteResponse{}, nil
}

// Update the routine in the database
func (s *DripsServer) RoutineUpdate(ctx context.Context, req *proto.RoutineUpdateRequest) (*proto.RoutineUpdateResponse, error) {
	_, err := s.db.Exec(`
	UPDATE routine
	SET name = ?, source = ?
	WHERE routine_id = ?
	`, req.Routine.Name, req.Routine.Source, req.Routine.RoutineId)
	if err != nil {
		return nil, err
	}

	var r proto.Routine
	err = s.db.QueryRow(`
	SELECT
		routine_id, name, source
	FROM routine
	WHERE routine_id = ?`, req.Routine.RoutineId).Scan(
		&r.RoutineId,
		&r.Name,
		&r.Source,
	)
	if err != nil {
		return nil, err
	}

	return &proto.RoutineUpdateResponse{
		Routine: &r,
	}, nil
}
