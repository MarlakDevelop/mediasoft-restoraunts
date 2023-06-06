package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"restaurant/internal/customer_office/repositories/officerepository"
	"restaurant/internal/customer_office/repositories/userrepository"
	customer "restaurant/pkg/contracts/customer_office"
	"restaurant/pkg/data/slice"
)

type UserServiceGRPCServer struct {
	userRepository   userrepository.UserRepository
	officeRepository officerepository.OfficeRepository
}

func NewUserServiceGRPCServer(userRepository userrepository.UserRepository, officeRepository officerepository.OfficeRepository) *UserServiceGRPCServer {
	return &UserServiceGRPCServer{
		userRepository:   userRepository,
		officeRepository: officeRepository,
	}
}

func (s *UserServiceGRPCServer) CreateUser(ctx context.Context, in *customer.CreateUserRequest) (*customer.CreateUserResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	_, err := s.officeRepository.Get(in.OfficeUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Office not found")
	}
	_, err = s.userRepository.Create(in.Name, in.OfficeUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &customer.CreateUserResponse{}, nil
}
func (s *UserServiceGRPCServer) GetUserList(ctx context.Context, in *customer.GetUserListRequest) (*customer.GetUserListResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	users, err := s.userRepository.List(in.OfficeUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &customer.GetUserListResponse{
		Result: slice.Map[*userrepository.User, *customer.User](users, func(index int, value *userrepository.User, slice []*userrepository.User) *customer.User {
			return &customer.User{
				Uuid:       value.Uuid,
				Name:       value.Name,
				OfficeUuid: value.Office.Uuid,
				OfficeName: value.Office.Name,
				CreatedAt:  timestamppb.New(value.CreatedAt),
			}
		}),
	}, nil
}
