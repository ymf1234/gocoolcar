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
type Service struct{}

func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	panic("")
}
