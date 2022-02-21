package profile

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type Fetcher interface {
	GetProfile(ctx context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error)
}

type Manager struct {
	Fetcher Fetcher
}

func (m *Manager) Verify(c context.Context, aid id.AccountID) (id.IdentityID, error) {
	nilID := id.IdentityID("")
	p, err := m.Fetcher.GetProfile(c, &rentalpb.GetProfileRequest{})

	if err != nil {
		return nilID, fmt.Errorf("cannot get profile: %v", err)
	}

	if p.IdentityStatus != rentalpb.IdentityStatus_VERIFIED {
		return nilID, fmt.Errorf("invalid identity status")
	}

	b, err := proto.Marshal(p.Identity)

	if err != nil {
		return nilID, fmt.Errorf("cannot marshal identity: %v", err)
	}

	return id.IdentityID(base64.StdEncoding.EncodeToString(b)), nil
}
