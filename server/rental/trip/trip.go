package trip

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
	rentalpb.UnimplementedTripServiceServer
}

func (s *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {

	return nil, status.Error(codes.Unimplemented, "")
}
