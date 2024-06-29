package dgrpc

import (
	"context"
	"database/sql"

	"github.com/mathyourlife/drips/proto"
)

// Implement your gRPC methods
func (s *DripsServer) ExerciseModifierCreate(ctx context.Context, req *proto.ExerciseModifierCreateRequest) (*proto.ExerciseModifierCreateResponse, error) {
	// Create the exercise modifier in the database

	result, err := s.db.Exec(`
        INSERT INTO exercise_modifier (exercise_id, modifier_id)
        VALUES (?, ?)
        `, req.ExerciseModifier.ExerciseId, req.ExerciseModifier.ModifierId)
	if err != nil {
		return nil, err
	}

	// parse the result
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Retrieve the inserted record
	var em proto.ExerciseModifier
	err = s.db.QueryRow(`
        SELECT
			exercise_modifier_id, exercise_id, modifier_id
        FROM exercise_modifier
        WHERE exercise_modifier_id = ?`, lastInsertID).Scan(
		&em.ExerciseModifierId,
		&em.ExerciseId,
		&em.ModifierId,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseModifierCreateResponse{
		ExerciseModifier: &em,
	}, nil
}

func (s *DripsServer) ExerciseModifier(ctx context.Context, req *proto.ExerciseModifierRequest) (*proto.ExerciseModifierResponse, error) {
	var em proto.ExerciseModifier
	err := s.db.QueryRow(`
	SELECT
		exercise_modifier_id, exercise_id, modifier_id
	FROM exercise_modifier
	WHERE exercise_modifier_id = ?`, req.ExerciseModifierId).Scan(
		&em.ExerciseModifierId,
		&em.ExerciseId,
		&em.ModifierId,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseModifierResponse{
		ExerciseModifier: &em,
	}, nil
}

// Implement your gRPC methods
func (s *DripsServer) ExerciseModifiers(ctx context.Context, req *proto.ExerciseModifiersRequest) (*proto.ExerciseModifiersResponse, error) {
	var ems []*proto.ExerciseModifier
	var rows *sql.Rows
	var err error

	if req.ExerciseId != 0 {
		// Filter by exercise ID
		rows, err = s.db.Query(`
				SELECT
					exercise_modifier_id, exercise_id, modifier_id
				FROM exercise_modifier
				WHERE exercise_id = ?
			`, req.ExerciseId)
	} else {
		// Return all exercise modifiers
		rows, err = s.db.Query(`
				SELECT
					exercise_modifier_id, exercise_id, modifier_id
				FROM exercise_modifier
			`)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var em proto.ExerciseModifier
		err := rows.Scan(&em.ExerciseModifierId, &em.ExerciseId, &em.ModifierId)
		if err != nil {
			return nil, err
		}
		ems = append(ems, &em)
	}

	return &proto.ExerciseModifiersResponse{
		ExerciseModifiers: ems,
	}, nil
}

// Delete the exercise modifier from the database
func (s *DripsServer) ExerciseModifierDelete(ctx context.Context, req *proto.ExerciseModifierDeleteRequest) (*proto.ExerciseModifierDeleteResponse, error) {
	_, err := s.db.Exec(`
        DELETE FROM exercise_modifier
        WHERE exercise_modifier_id = ?
        `, req.ExerciseModifierId)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseModifierDeleteResponse{}, nil
}

// Update the exercise modifier in the database
func (s *DripsServer) ExerciseModifierUpdate(ctx context.Context, req *proto.ExerciseModifierUpdateRequest) (*proto.ExerciseModifierUpdateResponse, error) {
	_, err := s.db.Exec(`
	UPDATE exercise_modifier
	SET exercise_id = ?, modifier_id = ?
	WHERE exercise_modifier_id = ?
	`, req.ExerciseModifier.ExerciseId, req.ExerciseModifier.ModifierId, req.ExerciseModifier.ExerciseModifierId)
	if err != nil {
		return nil, err
	}

	var em proto.ExerciseModifier
	err = s.db.QueryRow(`
	SELECT
		exercise_modifier_id, exercise_id, modifier_id
	FROM exercise_modifier
	WHERE exercise_modifier_id = ?`, req.ExerciseModifier.ExerciseModifierId).Scan(
		&em.ExerciseModifierId,
		&em.ExerciseId,
		&em.ModifierId,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ExerciseModifierUpdateResponse{
		ExerciseModifier: &em,
	}, nil
}
