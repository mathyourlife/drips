package dgrpc

import (
	"context"

	"github.com/mathyourlife/drips/proto"
)

// Implement your gRPC methods
func (s *DripsServer) ExerciseCreate(ctx context.Context, req *proto.ExerciseCreateRequest) (*proto.ExerciseCreateResponse, error) {
	// Create the exercise in the database

	result, err := s.db.Exec(`
        INSERT INTO exercise (exercise_class_id, name, duration_seconds, rest_seconds, repeat)
        VALUES (?, ?, ?, ?, ?)
        `, req.Exercise.ExerciseClassId, req.Exercise.Name, req.Exercise.DurationSeconds, req.Exercise.RestSeconds, req.Exercise.Repeat)
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
			exercise_id, name, exercise_class_id, duration_seconds, rest_seconds, repeat
        FROM exercise
        WHERE exercise_id = ?`, lastInsertID).Scan(
		&e.ExerciseId,
		&e.Name,
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

func (s *DripsServer) Exercise(ctx context.Context, req *proto.ExerciseRequest) (*proto.ExerciseResponse, error) {
	var e proto.Exercise
	err := s.db.QueryRow(`
	SELECT
		exercise_id, name, exercise_class_id, duration_seconds, rest_seconds, repeat
	FROM exercise
	WHERE exercise_id = ?`, req.ExerciseId).Scan(
		&e.ExerciseId,
		&e.Name,
		&e.ExerciseClassId,
		&e.DurationSeconds,
		&e.RestSeconds,
		&e.Repeat,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseResponse{
		Exercise: &e,
	}, nil
}

func (s *DripsServer) Exercises(ctx context.Context, req *proto.ExercisesRequest) (*proto.ExercisesResponse, error) {
	var es []*proto.Exercise
	rows, err := s.db.Query(`
        SELECT
			exercise_id, name, exercise_class_id, duration_seconds, rest_seconds, repeat
        FROM exercise`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var e proto.Exercise
		err := rows.Scan(&e.ExerciseId, &e.Name, &e.ExerciseClassId, &e.DurationSeconds, &e.RestSeconds, &e.Repeat)
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

// Update the exercise in the database
func (s *DripsServer) ExerciseUpdate(ctx context.Context, req *proto.ExerciseUpdateRequest) (*proto.ExerciseUpdateResponse, error) {
	_, err := s.db.Exec(`
	UPDATE exercise
	SET exercise_class_id = ?, name = ?, duration_seconds = ?, rest_seconds = ?, repeat = ?
	WHERE exercise_id = ?
	`, req.Exercise.ExerciseClassId, req.Exercise.Name, req.Exercise.DurationSeconds, req.Exercise.RestSeconds, req.Exercise.Repeat, req.Exercise.ExerciseId)
	if err != nil {
		return nil, err
	}

	var e proto.Exercise
	err = s.db.QueryRow(`
	SELECT
		exercise_id, name, exercise_class_id, duration_seconds, rest_seconds, repeat
	FROM exercise
	WHERE exercise_id = ?`, req.Exercise.ExerciseId).Scan(
		&e.ExerciseId,
		&e.Name,
		&e.ExerciseClassId,
		&e.DurationSeconds,
		&e.RestSeconds,
		&e.Repeat,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseUpdateResponse{
		Exercise: &e,
	}, nil
}
