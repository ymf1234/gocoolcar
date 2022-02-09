package profile

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/dao"

	"go.uber.org/zap"
)

type Service struct {
	Mongo  *dao.Mongo
	Logger *zap.Logger
	rentalpb.UnimplementedTripServiceServer
}

func (s *Service) GetProfile(ctx context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error) {
	return nil, nil
}

func (s *Service) SubmitProfile(ctx context.Context, req *rentalpb.Identity) (*rentalpb.Profile, error) {
	return nil, nil
}

func (s *Service) ClearProfile(ctx context.Context, req *rentalpb.ClearProfileRequest) (*rentalpb.Profile, error) {
	return nil, nil
}
