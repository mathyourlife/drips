package main

import (
	"github.com/mathyourlife/drips/proto"
)

// Define your gRPC server interface
type DripsServer struct {
	proto.UnimplementedDripsServiceServer
}
