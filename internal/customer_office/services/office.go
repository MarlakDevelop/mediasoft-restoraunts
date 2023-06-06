package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"restaurant/internal/customer_office/repositories/officerepository"
	customer "restaurant/pkg/contracts/customer_office"
	"restaurant/pkg/data/slice"
)

type OfficeServiceGRPCServer struct {
	officeRepository officerepository.OfficeRepository
}

func NewOfficeServiceGRPCServer(officeRepository officerepository.OfficeRepository) *OfficeServiceGRPCServer {
	return &OfficeServiceGRPCServer{
		officeRepository: officeRepository,
	}
}

func (s *OfficeServiceGRPCServer) CreateOffice(ctx context.Context, in *customer.CreateOfficeRequest) (*customer.CreateOfficeResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := s.officeRepository.Create(in.Name, in.Address)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &customer.CreateOfficeResponse{}, nil
}
func (s *OfficeServiceGRPCServer) GetOfficeList(ctx context.Context, in *customer.GetOfficeListRequest) (*customer.GetOfficeListResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	offices, err := s.officeRepository.List()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &customer.GetOfficeListResponse{
		Result: slice.Map[*officerepository.Office, *customer.Office](offices, func(index int, value *officerepository.Office, slice []*officerepository.Office) *customer.Office {
			return &customer.Office{
				Uuid:      value.Uuid,
				Name:      value.Name,
				Address:   value.Address,
				CreatedAt: timestamppb.New(value.CreatedAt),
			}
		}),
	}, nil
}
