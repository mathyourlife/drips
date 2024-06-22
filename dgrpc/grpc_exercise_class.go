package dgrpc

import (
	"context"

	"github.com/mathyourlife/drips/proto"
)

// Implement your gRPC methods
func (s *DripsServer) ExerciseClassCreate(ctx context.Context, req *proto.ExerciseClassCreateRequest) (*proto.ExerciseClassCreateResponse, error) {
	// Create the exercise class in the database

	result, err := s.db.Exec(`
	INSERT INTO exercise_class (name, short_name)
	VALUES (?, ?)
	`, req.ExerciseClass.Name, req.ExerciseClass.ShortName)
	if err != nil {
		return nil, err
	}

	// parse the result
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Retrieve the inserted record
	var ec proto.ExerciseClass
	err = s.db.QueryRow(`
	SELECT
		exercise_class_id, name, short_name
	FROM exercise_class
	WHERE exercise_class_id = ?`, lastInsertID).Scan(
		&ec.ExerciseClassId,
		&ec.Name,
		&ec.ShortName,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseClassCreateResponse{
		ExerciseClass: &ec,
	}, nil
}

// Implement your gRPC methods
func (s *DripsServer) ExerciseClasses(ctx context.Context, req *proto.ExerciseClassesRequest) (*proto.ExerciseClassesResponse, error) {

	// Return the list of exercise classes
	var ecs []*proto.ExerciseClass
	rows, err := s.db.Query(`
	SELECT
		exercise_class_id, name, short_name
	FROM exercise_class`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ec proto.ExerciseClass
		err := rows.Scan(&ec.ExerciseClassId, &ec.Name, &ec.ShortName)
		if err != nil {
			return nil, err
		}
		ecs = append(ecs, &ec)
	}

	return &proto.ExerciseClassesResponse{
		ExerciseClasses: ecs,
	}, nil
}

// Delete the exercise class from the database
func (s *DripsServer) ExerciseClassDelete(ctx context.Context, req *proto.ExerciseClassDeleteRequest) (*proto.ExerciseClassDeleteResponse, error) {
	_, err := s.db.Exec(`
	DELETE FROM exercise_class
	WHERE exercise_class_id = ?
	`, req.ExerciseClassId)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseClassDeleteResponse{}, nil
}
