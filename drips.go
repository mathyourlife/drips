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
