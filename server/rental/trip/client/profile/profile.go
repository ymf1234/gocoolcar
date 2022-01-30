package profile

import (
	"context"
	"coolcar/shared/id"
)

type Manager struct {
}

func (p *Manager) Verify(context.Context, id.AccountID) (id.IdentityID, error) {
	return id.IdentityID("identity1"), nil
}
