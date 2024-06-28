package dgrpc

import (
	"context"

	"github.com/mathyourlife/drips/proto"
)

// Implement your gRPC methods
func (s *DripsServer) ExerciseCreate(ctx context.Context, req *proto.ExerciseCreateRequest) (*proto.ExerciseCreateResponse, error) {
	// Create the exercise in the database

	result, err := s.db.Exec(`
        INSERT INTO exercise (name, exercise_class_id, duration_seconds, rest_seconds, repeat)
        VALUES (?, ?, ?, ?, ?)
        `, req.Exercise.Name, req.Exercise.ExerciseClassId, req.Exercise.DurationSeconds, req.Exercise.RestSeconds, req.Exercise.Repeat)
	if err != nil {
		return nil, err
	}

	// parse the result
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Retrieve the inserted record
	var e proto.Exercise
	err = s.db.QueryRow(`
        SELECT
			name, exercise_id, exercise_class_id, duration_seconds, rest_seconds, repeat
        FROM exercise
        WHERE exercise_id = ?`, lastInsertID).Scan(
		&e.Name,
		&e.ExerciseId,
		&e.ExerciseClassId,
		&e.DurationSeconds,
		&e.RestSeconds,
		&e.Repeat,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseCreateResponse{
		Exercise: &e,
	}, nil
}

func (s *DripsServer) Exercises(ctx context.Context, req *proto.ExercisesRequest) (*proto.ExercisesResponse, error) {
	var es []*proto.Exercise
	rows, err := s.db.Query(`
        SELECT
			name, exercise_id, exercise_class_id, duration_seconds, rest_seconds, repeat
        FROM exercise`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var e proto.Exercise
		err := rows.Scan(&e.Name, &e.ExerciseId, &e.ExerciseClassId, &e.DurationSeconds, &e.RestSeconds, &e.Repeat)
		if err != nil {
			return nil, err
		}
		es = append(es, &e)
	}

	return &proto.ExercisesResponse{
		Exercises: es,
	}, nil
}

// Delete the exercise from the database
func (s *DripsServer) ExerciseDelete(ctx context.Context, req *proto.ExerciseDeleteRequest) (*proto.ExerciseDeleteResponse, error) {
	_, err := s.db.Exec(`
        DELETE FROM exercise
        WHERE exercise_id = ?
        `, req.ExerciseId)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseDeleteResponse{}, nil
}
