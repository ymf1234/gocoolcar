package blob

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1"
	"coolcar/blob/dao"
	"coolcar/shared/id"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Mongo  *dao.Mongo
	Logger *zap.Logger
	blobpb.UnimplementedBlobServiceServer
}

func (s *Service) CreateBlob(c context.Context, req *blobpb.CreateBlobRequest) (*blobpb.CreateBlobResponse, error) {
	aid := id.AccountID(req.AccountId)
	br, err := s.Mongo.CreateBlob(c, aid)
	if err != nil {
		s.Logger.Error("cannot create blob", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) GetBlob(c context.Context, req *blobpb.GetBlobRequest) (*blobpb.GetBlobResponse, error) {
	return nil, nil
}

func (s *Service) GetBlobURL(c context.Context, req *blobpb.GetBlobURLRequest) (*blobpb.GetBlobURLResponse, error) {
	return nil, nil
}

func (s *Service) getBlobRecord(c context.Context, bid id.BlobID) (*dao.BlobRecord, error) {
	br, err := s.Mongo.GetBlob(c, bid)

	if err := res.Err(); err != nil {
		return nil, err
	}

	var br dao.BlobRecord
	if err := res.Decode(&br); err != nil {
		return nil, err
	}

	return &br, nil
}
