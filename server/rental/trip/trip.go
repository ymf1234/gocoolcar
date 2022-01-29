package trip

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	ProfileManager ProfileManager
	CarManager     CarManager
	POIManager     POIManager
	Mongo          *dao.Mongo
	Logger         *zap.Logger
	rentalpb.UnimplementedTripServiceServer
}

type ProfileManager interface {
	Verify(context.Context, id.AccountID) (id.IdentityID, error)
}

type CarManager interface {
	Verify(context.Context, id.CarID, *rentalpb.Location) error
	Unlock(context.Context, id.CarID) error
}

type POIManager interface {
	Resolve(context.Context, *rentalpb.Location) (string, error)
}

func (s *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// 验证驾驶者身份
	iID, err := s.ProfileManager.Verify(ctx, aid)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	// 检查车辆状态
	carID := id.CarID(req.CarId)
	err = s.CarManager.Verify(ctx, carID, req.Start)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	poi, err := s.POIManager.Resolve(ctx, req.Start)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("location ", req.Start), zap.Error(err))
	}
	// 创建行程: 写入数据库，开始计费
	ls := &rentalpb.LocationStatus{
		Location: req.Start,
		PoiName:  poi,
	}
	tr, err := s.Mongo.CreateTrip(ctx, &rentalpb.Trip{
		AccountId:  aid.String(),
		CarId:      carID.String(),
		IdentityId: iID.String(),
		Status:     rentalpb.TripStatus_IN_PROGRESS,
		Start:      ls,
		Current:    ls,
	})
	if err != nil {
		s.Logger.Warn("cannot create trip ", zap.Error(err))
		return nil, status.Error(codes.AlreadyExists, "")
	}

	// 车辆开锁
	go func() {
		err = s.CarManager.Unlock(context.Background(), carID)
		if err != nil {
			s.Logger.Error("cannot unlock car ", zap.Error(err))
		}
	}()

	return &rentalpb.TripEntity{
		Id:   tr.ID.Hex(),
		Trip: tr.Trip,
	}, nil
}

func (s *Service) GetTirp(ctx context.Context, req *rentalpb.GetTirpRequest) (*rentalpb.Trip, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) GetTirps(ctx context.Context, req *rentalpb.GetTirpsRequest) (*rentalpb.GetTirpsReponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) UpdateTrip(ctx context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	tid := id.TripID(req.Id)
	tr, err := s.Mongo.GetTirp(ctx, tid, aid)
	if req.Current != nil {
		tr.Trip.Current = s.calcCurrentStatus(tr.Trip, req.Current)
	}

	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
	}
	s.Mongo.UpdateTrip(ctx, tid, aid, tr.UpdateAt, tr.Trip)
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) calcCurrentStatus(trip *rentalpb.Trip, cur *rentalpb.Location) *rentalpb.LocationStatus {
	return nil
}
