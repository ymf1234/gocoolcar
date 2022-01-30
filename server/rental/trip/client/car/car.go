package car

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
)

type Manager struct {
}

func (c *Manager) Verify(context.Context, id.CarID, *rentalpb.Location) error {
	return nil
}

func (c *Manager) Unlock(context.Context, id.CarID) error {
	return nil
}
