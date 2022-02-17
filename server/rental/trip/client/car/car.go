package car

import (
	"context"
	carpb "coolcar/car/api/gen/v1"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
)

type Manager struct {
	CarService carpb.CarServiceClient
}

func (c *Manager) Verify(context.Context, id.CarID, *rentalpb.Location) error {
	return nil
}

func (c *Manager) Unlock(context.Context, id.CarID) error {
	return nil
}
