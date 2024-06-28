// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: proto/drips.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DripsServiceClient is the client API for DripsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DripsServiceClient interface {
	Exercise(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*ExerciseResponse, error)
	Exercises(ctx context.Context, in *ExercisesRequest, opts ...grpc.CallOption) (*ExercisesResponse, error)
	ExerciseCreate(ctx context.Context, in *ExerciseCreateRequest, opts ...grpc.CallOption) (*ExerciseCreateResponse, error)
	ExerciseDelete(ctx context.Context, in *ExerciseDeleteRequest, opts ...grpc.CallOption) (*ExerciseDeleteResponse, error)
	ExerciseClass(ctx context.Context, in *ExerciseClassRequest, opts ...grpc.CallOption) (*ExerciseClassResponse, error)
	ExerciseClasses(ctx context.Context, in *ExerciseClassesRequest, opts ...grpc.CallOption) (*ExerciseClassesResponse, error)
	ExerciseClassCreate(ctx context.Context, in *ExerciseClassCreateRequest, opts ...grpc.CallOption) (*ExerciseClassCreateResponse, error)
	ExerciseClassDelete(ctx context.Context, in *ExerciseClassDeleteRequest, opts ...grpc.CallOption) (*ExerciseClassDeleteResponse, error)
	ExerciseModifier(ctx context.Context, in *ExerciseModifierRequest, opts ...grpc.CallOption) (*ExerciseModifierResponse, error)
	ExerciseModifiers(ctx context.Context, in *ExerciseModifiersRequest, opts ...grpc.CallOption) (*ExerciseModifiersResponse, error)
	ExerciseModifierCreate(ctx context.Context, in *ExerciseModifierCreateRequest, opts ...grpc.CallOption) (*ExerciseModifierCreateResponse, error)
	ExerciseModifierDelete(ctx context.Context, in *ExerciseModifierDeleteRequest, opts ...grpc.CallOption) (*ExerciseModifierDeleteResponse, error)
	Modifier(ctx context.Context, in *ModifierRequest, opts ...grpc.CallOption) (*ModifierResponse, error)
	Modifiers(ctx context.Context, in *ModifiersRequest, opts ...grpc.CallOption) (*ModifiersResponse, error)
	ModifierCreate(ctx context.Context, in *ModifierCreateRequest, opts ...grpc.CallOption) (*ModifierCreateResponse, error)
	ModifierDelete(ctx context.Context, in *ModifierDeleteRequest, opts ...grpc.CallOption) (*ModifierDeleteResponse, error)
	Routine(ctx context.Context, in *RoutineRequest, opts ...grpc.CallOption) (*RoutineResponse, error)
	Routines(ctx context.Context, in *RoutinesRequest, opts ...grpc.CallOption) (*RoutinesResponse, error)
	RoutineCreate(ctx context.Context, in *RoutineCreateRequest, opts ...grpc.CallOption) (*RoutineCreateResponse, error)
	RoutineDelete(ctx context.Context, in *RoutineDeleteRequest, opts ...grpc.CallOption) (*RoutineDeleteResponse, error)
	User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Users(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error)
}

type dripsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDripsServiceClient(cc grpc.ClientConnInterface) DripsServiceClient {
	return &dripsServiceClient{cc}
}

func (c *dripsServiceClient) Exercise(ctx context.Context, in *ExerciseRequest, opts ...grpc.CallOption) (*ExerciseResponse, error) {
	out := new(ExerciseResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/Exercise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) Exercises(ctx context.Context, in *ExercisesRequest, opts ...grpc.CallOption) (*ExercisesResponse, error) {
	out := new(ExercisesResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/Exercises", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseCreate(ctx context.Context, in *ExerciseCreateRequest, opts ...grpc.CallOption) (*ExerciseCreateResponse, error) {
	out := new(ExerciseCreateResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseDelete(ctx context.Context, in *ExerciseDeleteRequest, opts ...grpc.CallOption) (*ExerciseDeleteResponse, error) {
	out := new(ExerciseDeleteResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseClass(ctx context.Context, in *ExerciseClassRequest, opts ...grpc.CallOption) (*ExerciseClassResponse, error) {
	out := new(ExerciseClassResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseClasses(ctx context.Context, in *ExerciseClassesRequest, opts ...grpc.CallOption) (*ExerciseClassesResponse, error) {
	out := new(ExerciseClassesResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseClasses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseClassCreate(ctx context.Context, in *ExerciseClassCreateRequest, opts ...grpc.CallOption) (*ExerciseClassCreateResponse, error) {
	out := new(ExerciseClassCreateResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseClassCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseClassDelete(ctx context.Context, in *ExerciseClassDeleteRequest, opts ...grpc.CallOption) (*ExerciseClassDeleteResponse, error) {
	out := new(ExerciseClassDeleteResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseClassDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseModifier(ctx context.Context, in *ExerciseModifierRequest, opts ...grpc.CallOption) (*ExerciseModifierResponse, error) {
	out := new(ExerciseModifierResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseModifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseModifiers(ctx context.Context, in *ExerciseModifiersRequest, opts ...grpc.CallOption) (*ExerciseModifiersResponse, error) {
	out := new(ExerciseModifiersResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseModifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseModifierCreate(ctx context.Context, in *ExerciseModifierCreateRequest, opts ...grpc.CallOption) (*ExerciseModifierCreateResponse, error) {
	out := new(ExerciseModifierCreateResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseModifierCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ExerciseModifierDelete(ctx context.Context, in *ExerciseModifierDeleteRequest, opts ...grpc.CallOption) (*ExerciseModifierDeleteResponse, error) {
	out := new(ExerciseModifierDeleteResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ExerciseModifierDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) Modifier(ctx context.Context, in *ModifierRequest, opts ...grpc.CallOption) (*ModifierResponse, error) {
	out := new(ModifierResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/Modifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) Modifiers(ctx context.Context, in *ModifiersRequest, opts ...grpc.CallOption) (*ModifiersResponse, error) {
	out := new(ModifiersResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/Modifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ModifierCreate(ctx context.Context, in *ModifierCreateRequest, opts ...grpc.CallOption) (*ModifierCreateResponse, error) {
	out := new(ModifierCreateResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ModifierCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) ModifierDelete(ctx context.Context, in *ModifierDeleteRequest, opts ...grpc.CallOption) (*ModifierDeleteResponse, error) {
	out := new(ModifierDeleteResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/ModifierDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) Routine(ctx context.Context, in *RoutineRequest, opts ...grpc.CallOption) (*RoutineResponse, error) {
	out := new(RoutineResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/Routine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) Routines(ctx context.Context, in *RoutinesRequest, opts ...grpc.CallOption) (*RoutinesResponse, error) {
	out := new(RoutinesResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/Routines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) RoutineCreate(ctx context.Context, in *RoutineCreateRequest, opts ...grpc.CallOption) (*RoutineCreateResponse, error) {
	out := new(RoutineCreateResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/RoutineCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) RoutineDelete(ctx context.Context, in *RoutineDeleteRequest, opts ...grpc.CallOption) (*RoutineDeleteResponse, error) {
	out := new(RoutineDeleteResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/RoutineDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) Users(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/Users", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) UserCreate(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/UserCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dripsServiceClient) UserDelete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserDeleteResponse, error) {
	out := new(UserDeleteResponse)
	err := c.cc.Invoke(ctx, "/drips.DripsService/UserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DripsServiceServer is the server API for DripsService service.
// All implementations must embed UnimplementedDripsServiceServer
// for forward compatibility
type DripsServiceServer interface {
	Exercise(context.Context, *ExerciseRequest) (*ExerciseResponse, error)
	Exercises(context.Context, *ExercisesRequest) (*ExercisesResponse, error)
	ExerciseCreate(context.Context, *ExerciseCreateRequest) (*ExerciseCreateResponse, error)
	ExerciseDelete(context.Context, *ExerciseDeleteRequest) (*ExerciseDeleteResponse, error)
	ExerciseClass(context.Context, *ExerciseClassRequest) (*ExerciseClassResponse, error)
	ExerciseClasses(context.Context, *ExerciseClassesRequest) (*ExerciseClassesResponse, error)
	ExerciseClassCreate(context.Context, *ExerciseClassCreateRequest) (*ExerciseClassCreateResponse, error)
	ExerciseClassDelete(context.Context, *ExerciseClassDeleteRequest) (*ExerciseClassDeleteResponse, error)
	ExerciseModifier(context.Context, *ExerciseModifierRequest) (*ExerciseModifierResponse, error)
	ExerciseModifiers(context.Context, *ExerciseModifiersRequest) (*ExerciseModifiersResponse, error)
	ExerciseModifierCreate(context.Context, *ExerciseModifierCreateRequest) (*ExerciseModifierCreateResponse, error)
	ExerciseModifierDelete(context.Context, *ExerciseModifierDeleteRequest) (*ExerciseModifierDeleteResponse, error)
	Modifier(context.Context, *ModifierRequest) (*ModifierResponse, error)
	Modifiers(context.Context, *ModifiersRequest) (*ModifiersResponse, error)
	ModifierCreate(context.Context, *ModifierCreateRequest) (*ModifierCreateResponse, error)
	ModifierDelete(context.Context, *ModifierDeleteRequest) (*ModifierDeleteResponse, error)
	Routine(context.Context, *RoutineRequest) (*RoutineResponse, error)
	Routines(context.Context, *RoutinesRequest) (*RoutinesResponse, error)
	RoutineCreate(context.Context, *RoutineCreateRequest) (*RoutineCreateResponse, error)
	RoutineDelete(context.Context, *RoutineDeleteRequest) (*RoutineDeleteResponse, error)
	User(context.Context, *UserRequest) (*UserResponse, error)
	Users(context.Context, *UsersRequest) (*UsersResponse, error)
	UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error)
	mustEmbedUnimplementedDripsServiceServer()
}

// UnimplementedDripsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDripsServiceServer struct {
}

func (UnimplementedDripsServiceServer) Exercise(context.Context, *ExerciseRequest) (*ExerciseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exercise not implemented")
}
func (UnimplementedDripsServiceServer) Exercises(context.Context, *ExercisesRequest) (*ExercisesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exercises not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseCreate(context.Context, *ExerciseCreateRequest) (*ExerciseCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseCreate not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseDelete(context.Context, *ExerciseDeleteRequest) (*ExerciseDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseDelete not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseClass(context.Context, *ExerciseClassRequest) (*ExerciseClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseClass not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseClasses(context.Context, *ExerciseClassesRequest) (*ExerciseClassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseClasses not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseClassCreate(context.Context, *ExerciseClassCreateRequest) (*ExerciseClassCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseClassCreate not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseClassDelete(context.Context, *ExerciseClassDeleteRequest) (*ExerciseClassDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseClassDelete not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseModifier(context.Context, *ExerciseModifierRequest) (*ExerciseModifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseModifier not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseModifiers(context.Context, *ExerciseModifiersRequest) (*ExerciseModifiersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseModifiers not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseModifierCreate(context.Context, *ExerciseModifierCreateRequest) (*ExerciseModifierCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseModifierCreate not implemented")
}
func (UnimplementedDripsServiceServer) ExerciseModifierDelete(context.Context, *ExerciseModifierDeleteRequest) (*ExerciseModifierDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExerciseModifierDelete not implemented")
}
func (UnimplementedDripsServiceServer) Modifier(context.Context, *ModifierRequest) (*ModifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modifier not implemented")
}
func (UnimplementedDripsServiceServer) Modifiers(context.Context, *ModifiersRequest) (*ModifiersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modifiers not implemented")
}
func (UnimplementedDripsServiceServer) ModifierCreate(context.Context, *ModifierCreateRequest) (*ModifierCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifierCreate not implemented")
}
func (UnimplementedDripsServiceServer) ModifierDelete(context.Context, *ModifierDeleteRequest) (*ModifierDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifierDelete not implemented")
}
func (UnimplementedDripsServiceServer) Routine(context.Context, *RoutineRequest) (*RoutineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Routine not implemented")
}
func (UnimplementedDripsServiceServer) Routines(context.Context, *RoutinesRequest) (*RoutinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Routines not implemented")
}
func (UnimplementedDripsServiceServer) RoutineCreate(context.Context, *RoutineCreateRequest) (*RoutineCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoutineCreate not implemented")
}
func (UnimplementedDripsServiceServer) RoutineDelete(context.Context, *RoutineDeleteRequest) (*RoutineDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoutineDelete not implemented")
}
func (UnimplementedDripsServiceServer) User(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedDripsServiceServer) Users(context.Context, *UsersRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Users not implemented")
}
func (UnimplementedDripsServiceServer) UserCreate(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreate not implemented")
}
func (UnimplementedDripsServiceServer) UserDelete(context.Context, *UserDeleteRequest) (*UserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (UnimplementedDripsServiceServer) mustEmbedUnimplementedDripsServiceServer() {}

// UnsafeDripsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DripsServiceServer will
// result in compilation errors.
type UnsafeDripsServiceServer interface {
	mustEmbedUnimplementedDripsServiceServer()
}

func RegisterDripsServiceServer(s grpc.ServiceRegistrar, srv DripsServiceServer) {
	s.RegisterService(&DripsService_ServiceDesc, srv)
}

func _DripsService_Exercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).Exercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/Exercise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).Exercise(ctx, req.(*ExerciseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_Exercises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExercisesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).Exercises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/Exercises",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).Exercises(ctx, req.(*ExercisesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseCreate(ctx, req.(*ExerciseCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseDelete(ctx, req.(*ExerciseDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseClass(ctx, req.(*ExerciseClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseClassesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseClasses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseClasses(ctx, req.(*ExerciseClassesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseClassCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseClassCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseClassCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseClassCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseClassCreate(ctx, req.(*ExerciseClassCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseClassDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseClassDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseClassDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseClassDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseClassDelete(ctx, req.(*ExerciseClassDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseModifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseModifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseModifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseModifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseModifier(ctx, req.(*ExerciseModifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseModifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseModifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseModifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseModifiers(ctx, req.(*ExerciseModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseModifierCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseModifierCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseModifierCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseModifierCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseModifierCreate(ctx, req.(*ExerciseModifierCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ExerciseModifierDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExerciseModifierDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ExerciseModifierDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ExerciseModifierDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ExerciseModifierDelete(ctx, req.(*ExerciseModifierDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_Modifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).Modifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/Modifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).Modifier(ctx, req.(*ModifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_Modifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).Modifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/Modifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).Modifiers(ctx, req.(*ModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ModifierCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifierCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ModifierCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ModifierCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ModifierCreate(ctx, req.(*ModifierCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_ModifierDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifierDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).ModifierDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/ModifierDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).ModifierDelete(ctx, req.(*ModifierDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_Routine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).Routine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/Routine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).Routine(ctx, req.(*RoutineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_Routines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).Routines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/Routines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).Routines(ctx, req.(*RoutinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_RoutineCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutineCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).RoutineCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/RoutineCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).RoutineCreate(ctx, req.(*RoutineCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_RoutineDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutineDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).RoutineDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/RoutineDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).RoutineDelete(ctx, req.(*RoutineDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).User(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_Users_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).Users(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/Users",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).Users(ctx, req.(*UsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_UserCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).UserCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/UserCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).UserCreate(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DripsService_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DripsServiceServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drips.DripsService/UserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DripsServiceServer).UserDelete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DripsService_ServiceDesc is the grpc.ServiceDesc for DripsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DripsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "drips.DripsService",
	HandlerType: (*DripsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exercise",
			Handler:    _DripsService_Exercise_Handler,
		},
		{
			MethodName: "Exercises",
			Handler:    _DripsService_Exercises_Handler,
		},
		{
			MethodName: "ExerciseCreate",
			Handler:    _DripsService_ExerciseCreate_Handler,
		},
		{
			MethodName: "ExerciseDelete",
			Handler:    _DripsService_ExerciseDelete_Handler,
		},
		{
			MethodName: "ExerciseClass",
			Handler:    _DripsService_ExerciseClass_Handler,
		},
		{
			MethodName: "ExerciseClasses",
			Handler:    _DripsService_ExerciseClasses_Handler,
		},
		{
			MethodName: "ExerciseClassCreate",
			Handler:    _DripsService_ExerciseClassCreate_Handler,
		},
		{
			MethodName: "ExerciseClassDelete",
			Handler:    _DripsService_ExerciseClassDelete_Handler,
		},
		{
			MethodName: "ExerciseModifier",
			Handler:    _DripsService_ExerciseModifier_Handler,
		},
		{
			MethodName: "ExerciseModifiers",
			Handler:    _DripsService_ExerciseModifiers_Handler,
		},
		{
			MethodName: "ExerciseModifierCreate",
			Handler:    _DripsService_ExerciseModifierCreate_Handler,
		},
		{
			MethodName: "ExerciseModifierDelete",
			Handler:    _DripsService_ExerciseModifierDelete_Handler,
		},
		{
			MethodName: "Modifier",
			Handler:    _DripsService_Modifier_Handler,
		},
		{
			MethodName: "Modifiers",
			Handler:    _DripsService_Modifiers_Handler,
		},
		{
			MethodName: "ModifierCreate",
			Handler:    _DripsService_ModifierCreate_Handler,
		},
		{
			MethodName: "ModifierDelete",
			Handler:    _DripsService_ModifierDelete_Handler,
		},
		{
			MethodName: "Routine",
			Handler:    _DripsService_Routine_Handler,
		},
		{
			MethodName: "Routines",
			Handler:    _DripsService_Routines_Handler,
		},
		{
			MethodName: "RoutineCreate",
			Handler:    _DripsService_RoutineCreate_Handler,
		},
		{
			MethodName: "RoutineDelete",
			Handler:    _DripsService_RoutineDelete_Handler,
		},
		{
			MethodName: "User",
			Handler:    _DripsService_User_Handler,
		},
		{
			MethodName: "Users",
			Handler:    _DripsService_Users_Handler,
		},
		{
			MethodName: "UserCreate",
			Handler:    _DripsService_UserCreate_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _DripsService_UserDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/drips.proto",
}
