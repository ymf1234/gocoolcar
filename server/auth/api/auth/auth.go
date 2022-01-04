package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	OpenIDResolver OpenIDResolver
	authpb.UnimplementedAuthServiceServer
	Logger *zap.Logger
}

type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	openID, err := s.OpenIDResolver.Resolve(req.Code)

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: &v", err)
	}
	s.Logger.Info("received code ", zap.String("code", req.Code))

	return &authpb.LoginResponse{
		AccessToken: "token for" + openID,
		ExpiresIn:   7200,
	}, nil
}
