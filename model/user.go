package model

import (
	"gorm.io/gorm"

	pb "github.com/mathyourlife/drips/proto"
)

type User struct {
	gorm.Model
	DisplayName string `gorm:"uniqueIndex"`
}

func NewUserFromProto(u *pb.User) User {
	user := User{
		DisplayName: u.GetDisplayName(),
	}
	user.ID = uint(u.GetUserId())
	return user
}

func (m User) String() string {
	return m.DisplayName
}

func (m User) ToProto() *pb.User {
	return &pb.User{
		UserId:      int32(m.ID),
		DisplayName: m.DisplayName,
	}
}
