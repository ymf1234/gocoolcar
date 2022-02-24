package blob

import (
	"context"
	"coolcar/auth/dao"
	blobpb "coolcar/blob/api/gen/v1"

	"go.uber.org/zap"
)

type Service struct {
	Mongo  *dao.Mongo
	Logger *zap.Logger
	blobpb.UnimplementedBlobServiceServer
}


func (s *Service) CreateBlob(c context.Context, blobpb.CreateBlobRequest) (*blobpb.CreateBlobResponse, error) {
	return nil, nil
}