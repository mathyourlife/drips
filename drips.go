package drips

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mathyourlife/drips/model"
	pb "github.com/mathyourlife/drips/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Service struct {
	pb.UnimplementedDripsServiceServer
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Exercise(ctx context.Context, req *pb.ExerciseRequest) (*pb.ExerciseResponse, error) {
	var e model.Exercise
	s.db.First(&e, req.ExerciseId)

	return &pb.ExerciseResponse{
		Exercise: e.ToProto(),
	}, nil
}

func (s *Service) Exercises(ctx context.Context, req *pb.ExercisesRequest) (*pb.ExercisesResponse, error) {
	var exercises []model.Exercise
	s.db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name)).Find(&exercises)

	var es []*pb.Exercise
	for _, e := range exercises {
		es = append(es, e.ToProto())
	}

	return &pb.ExercisesResponse{
		Exercises: es,
	}, nil
}

func (s *Service) ExerciseCreate(ctx context.Context, req *pb.ExerciseCreateRequest) (*pb.ExerciseCreateResponse, error) {
	exercise := req.Exercise

	e := model.NewExerciseFromProto(exercise)
	if err := s.db.Create(&e).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to create exercise: %s", err)
	}
	return &pb.ExerciseCreateResponse{
		Exercise: e.ToProto(),
	}, nil
}

func (s *Service) ExerciseDelete(ctx context.Context, req *pb.ExerciseDeleteRequest) (*pb.ExerciseDeleteResponse, error) {
	if err := s.db.Delete(&model.Exercise{}, req.ExerciseId).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to delete user-id %d: %s", req.ExerciseId, err)
	}
	return &pb.ExerciseDeleteResponse{}, nil
}

func (s *Service) ExerciseClass(ctx context.Context, req *pb.ExerciseClassRequest) (*pb.ExerciseClassResponse, error) {
	var u model.ExerciseClass
	s.db.First(&u, req.ExerciseClassId)

	return &pb.ExerciseClassResponse{
		ExerciseClass: u.ToProto(),
	}, nil
}

func (s *Service) ExerciseClasses(ctx context.Context, req *pb.ExerciseClassesRequest) (*pb.ExerciseClassesResponse, error) {
	var users []model.ExerciseClass
	s.db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name)).Find(&users)

	var us []*pb.ExerciseClass
	for _, u := range users {
		us = append(us, u.ToProto())
	}

	return &pb.ExerciseClassesResponse{
		ExerciseClasses: us,
	}, nil
}

func (s *Service) ExerciseClassCreate(ctx context.Context, req *pb.ExerciseClassCreateRequest) (*pb.ExerciseClassCreateResponse, error) {
	user := req.ExerciseClass

	if user.Name == "" {
		return nil, status.Errorf(codes.Internal, "ExerciseClass requires a Name")
	}

	u := model.NewExerciseClassFromProto(user)
	if err := s.db.Create(&u).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to create exercise class: %s", err)
	}
	return &pb.ExerciseClassCreateResponse{
		ExerciseClass: u.ToProto(),
	}, nil
}

func (s *Service) ExerciseClassDelete(ctx context.Context, req *pb.ExerciseClassDeleteRequest) (*pb.ExerciseClassDeleteResponse, error) {
	if err := s.db.Delete(&model.ExerciseClass{}, req.ExerciseClassId).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to delete user-id %d: %s", req.ExerciseClassId, err)
	}
	return &pb.ExerciseClassDeleteResponse{}, nil
}

func (s *Service) Modifier(ctx context.Context, req *pb.ModifierRequest) (*pb.ModifierResponse, error) {
	var modifier model.Modifier
	if err := s.db.First(&modifier, req.ModiferId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.InvalidArgument, "Unable to locate modifier-id %d", req.ModiferId)
	}

	return &pb.ModifierResponse{
		Modifier: modifier.ToProto(),
	}, nil
}

func (s *Service) Modifiers(ctx context.Context, req *pb.ModifiersRequest) (*pb.ModifiersResponse, error) {
	var modifiers []model.Modifier
	s.db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name)).Find(&modifiers)

	var mods []*pb.Modifier
	for _, m := range modifiers {
		mods = append(mods, m.ToProto())
	}

	return &pb.ModifiersResponse{
		Modifiers: mods,
	}, nil
}

func (s *Service) ModifierCreate(ctx context.Context, req *pb.ModifierCreateRequest) (*pb.ModifierCreateResponse, error) {
	modifier := req.Modifier

	if modifier.Name == "" {
		return nil, status.Errorf(codes.Internal, "Modifier requires a Name")
	}

	mod := model.NewModifierFromProto(modifier)
	if err := s.db.Create(&mod).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to create modifier: %s", err)
	}
	return &pb.ModifierCreateResponse{
		Modifier: modifier,
	}, nil
}

func (s *Service) ModifierDelete(ctx context.Context, req *pb.ModifierDeleteRequest) (*pb.ModifierDeleteResponse, error) {
	if err := s.db.Delete(&model.Modifier{}, req.ModifierId).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to delete modifier-id %d: %s", req.ModifierId, err)
	}
	return &pb.ModifierDeleteResponse{}, nil
}

func (s *Service) User(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	var u model.User
	s.db.First(&u, req.UserId)

	return &pb.UserResponse{
		User: u.ToProto(),
	}, nil
}

func (s *Service) Users(ctx context.Context, req *pb.UsersRequest) (*pb.UsersResponse, error) {
	var users []model.User
	s.db.Where("display_name LIKE ?", fmt.Sprintf("%%%s%%", req.DisplayName)).Find(&users)

	var us []*pb.User
	for _, u := range users {
		us = append(us, u.ToProto())
	}

	return &pb.UsersResponse{
		Users: us,
	}, nil
}

func (s *Service) UserCreate(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	user := req.User

	if user.DisplayName == "" {
		return nil, status.Errorf(codes.Internal, "User requires a DisplayName")
	}

	u := model.NewUserFromProto(user)
	if err := s.db.Create(&u).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to create modifier: %s", err)
	}
	return &pb.UserCreateResponse{
		User: u.ToProto(),
	}, nil
}

func (s *Service) UserDelete(ctx context.Context, req *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	if err := s.db.Delete(&model.User{}, req.UserId).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to delete user-id %d: %s", req.UserId, err)
	}
	return &pb.UserDeleteResponse{}, nil
}

func (s *Service) Routine(ctx context.Context, req *pb.RoutineRequest) (*pb.RoutineResponse, error) {
	// // Load the routine
	// query := `SELECT routine_id, name, source, sequence
	// 	FROM routine
	// 	WHERE routine_id = $1;`

	// r := &pb.Routine{}
	// err := s.db.QueryRowContext(ctx, query, req.RoutineId).Scan(
	// 	&r.RoutineId, &r.Name, &r.Source, &r.Sequence)
	// if err != nil {
	// 	return nil, err
	// }

	// // Load exercises
	// query = `SELECT exercise.exercise_id, exercise.sequence, class.name, exercise.duration_seconds, exercise.rest_seconds
	// 	FROM routine_exercise
	// 	JOIN exercise ON routine_exercise.exercise_id = exercise.exercise_id
	// 	JOIN class ON exercise.class_id = class.class_id
	// 	WHERE routine_exercise.routine_id = $1`

	// rows, err := s.db.Query(query, req.RoutineId)
	// if err != nil {
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	e := &pb.Exercise{
	// 		Class: &pb.ExerciseClass{},
	// 	}
	// 	var durationSec, restSec int
	// 	if err := rows.Scan(&e.ExerciseId, &e.Sequence, &e.Class.Name, &durationSec, &restSec); err != nil {
	// 		return nil, err
	// 	}
	// 	e.Duration = durationpb.New(time.Duration(durationSec) * time.Second)
	// 	e.Rest = durationpb.New(time.Duration(restSec) * time.Second)
	// 	r.Exercises = append(r.Exercises, e)
	// }

	// // Load modifiers for the exercise
	// query = `SELECT modifier.name
	// 	FROM exercise_modifier
	// 	JOIN modifier ON exercise_modifier.modifier_id = modifier.modifier_id
	// 	WHERE exercise_modifier.exercise_id = $1`

	// for _, e := range r.Exercises {
	// 	rows, err = s.db.Query(query, e.ExerciseId)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	for rows.Next() {
	// 		var modifier string
	// 		if err := rows.Scan(&modifier); err != nil {
	// 			return nil, err
	// 		}
	// 		e.Modifiers = append(e.Modifiers, &pb.Modifier{Name: modifier})
	// 	}
	// }

	return &pb.RoutineResponse{
		// Routine: r,
	}, nil
}

func (s *Service) Routines(ctx context.Context, req *pb.RoutinesRequest) (*pb.RoutinesResponse, error) {

	// // Search exercises
	// query := `SELECT routine_id, name, source, sequence
	// 	FROM routine;`

	// rows, err := s.db.Query(query)
	// if err != nil {
	// 	return nil, err
	// }

	// var rs []*pb.Routine
	// for rows.Next() {
	// 	r := &pb.Routine{}
	// 	if err := rows.Scan(&r.RoutineId, &r.Name, &r.Source, &r.Sequence); err != nil {
	// 		return nil, err
	// 	}
	// 	// Match for a substring of the routine name or match all if search name is empty.
	// 	if strings.Contains(strings.ToLower(r.Name), strings.ToLower(req.Name)) {
	// 		rs = append(rs, r)
	// 	}
	// }

	return &pb.RoutinesResponse{
		// Routines: rs,
	}, nil
}

func PrintRoutine(routine *pb.Routine) string {
	var exs []string
	for _, e := range routine.Exercises {
		var l string
		var modStrs []string
		for _, mod := range e.Modifiers {
			modStrs = append(modStrs, mod.Name)
		}
		l = fmt.Sprintf("%d: %s (%s) for %d seconds", e.Sequence+1, e.Class.Name, strings.Join(modStrs, ","), e.Duration.Seconds)
		if e.Rest != nil && e.Rest.AsDuration() != 0 {
			l += fmt.Sprintf(" then rest for %d seconds", e.Rest.Seconds)
		}

		exs = append(exs, l)
	}

	out := fmt.Sprintf("%s\n#%d\n\n%s", routine.Name, routine.Sequence+1, strings.Join(exs, "\n"))
	return out
}
