package dgrpc

import (
	"database/sql"

	"github.com/mathyourlife/drips/proto"
)

// Define your gRPC server interface
type DripsServer struct {
	db *sql.DB
	proto.UnimplementedDripsServiceServer
}

func NewServer(db *sql.DB) *DripsServer {
	return &DripsServer{
		db: db,
	}
}
