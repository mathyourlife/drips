package drips

import (
	"context"

	pb "github.com/mathyourlife/drips/proto"
)

type Service struct {
	pb.UnimplementedDripsServiceServer
}

func (s *Service) User(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		User: &pb.User{
			UserId:      2,
			DisplayName: "Dan",
		},
	}, nil
}

func (s *Service) Routine(ctx context.Context, req *pb.RoutineRequest) (*pb.RoutineResponse, error) {
	return &pb.RoutineResponse{
		Routine: &pb.Routine{
			Name: "Caroline Girvan",
		},
	}, nil
}
