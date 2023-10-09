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
	db := s.db.Joins("ExerciseClass").Preload("Modifiers")
	if req.ExerciseClassId > 0 {
		db = db.Where("exercise_class_id = ?", req.ExerciseClassId)
	}
	db.Find(&exercises)

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
	var ec model.ExerciseClass
	if req.ExerciseClassId > 0 {
		s.db.First(&ec, req.ExerciseClassId)
	} else if req.Name != "" {
		s.db.First(&ec, "name = ?", req.Name)
	}

	return &pb.ExerciseClassResponse{
		ExerciseClass: ec.ToProto(),
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
	var r model.Routine
	db := s.db.Preload("Exercises")
	db.First(&r, req.RoutineId)

	return &pb.RoutineResponse{
		Routine: r.ToProto(),
	}, nil
}

func (s *Service) Routines(ctx context.Context, req *pb.RoutinesRequest) (*pb.RoutinesResponse, error) {
	var routines []model.Routine

	db := s.db.Preload("Exercises")
	if req.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Name))
	}
	db.Find(&routines)

	var rs []*pb.Routine
	for _, r := range routines {
		rs = append(rs, r.ToProto())
	}

	return &pb.RoutinesResponse{
		Routines: rs,
	}, nil
}

func (s *Service) RoutineCreate(ctx context.Context, req *pb.RoutineCreateRequest) (*pb.RoutineCreateResponse, error) {
	routine := req.Routine

	r := model.NewRoutineFromProto(routine)
	if err := s.db.Create(&r).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to create routine: %s", err)
	}
	return &pb.RoutineCreateResponse{
		Routine: r.ToProto(),
	}, nil
}

func (s *Service) RoutineDelete(ctx context.Context, req *pb.RoutineDeleteRequest) (*pb.RoutineDeleteResponse, error) {
	if err := s.db.Delete(&model.Routine{}, req.RoutineId).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to delete routine-id %d: %s", req.RoutineId, err)
	}
	return &pb.RoutineDeleteResponse{}, nil
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
