package poi

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"hash/fnv"

	"google.golang.org/protobuf/proto"
)

var poi = []string{
	"中关村",
	"天安门",
	"陆家嘴",
	"迪士尼",
	"天河",
	"世纪大道",
}

type Manager struct {
}

func (*Manager) Resolve(c context.Context, req *rentalpb.Location) (string, error) {
	b, err := proto.Marshal(req)
	if err != nil {
		return "", err
	}
	h := fnv.New32()
	h.Write(b)
	return poi[int(h.Sum32())%len(poi)], nil
}
