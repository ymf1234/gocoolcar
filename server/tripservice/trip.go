package trip

import (
	"context"
	trippb "coolcar/proto/gen/go"
)

// type TripServiceServer interface {
// 	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
// 	mustEmbedUnimplementedTripServiceServer()
// }

// Service implements trip service.
type Service struct {
	trippb.UnimplementedTripServiceServer
}

func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "asc",
			End:         "fff",
			DurationSec: 10,
			FeeCent:     10,
			StartPos: &trippb.Location{
				Latitude:  33.11,
				Longitude: 120,
			},
			EndPos: &trippb.Location{
				Latitude:  35,
				Longitude: 115,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  35,
					Longitude: 115,
				},
				{
					Latitude:  32,
					Longitude: 105,
				},
			},
			Status:          trippb.TripStatus_FINISHED,
			IsPromotionTrip: true,
		},
	}, nil
}
