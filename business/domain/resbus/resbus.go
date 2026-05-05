package resbus

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotFound    = errors.New("reservation not found")
	ErrNoAvailable = errors.New("reservation not available to be completed")
)

// ExtBusiness interface provides support for extensions that wrap extra functionality
// around the core business logic.
type ExtBusiness interface {
	NewWithTx(tx sqldb.CommitRollbacker) (ExtBusiness, error)
	Create(ctx context.Context, np NewProduct) (Product, error)
	Update(ctx context.Context, prd Reservation, up UpdateProduct) (Product, error)
	Delete(ctx context.Context, prd Reservation) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, page page.Page) ([]Product, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, productID uuid.UUID) (Reservation, error)
	QueryByUserID(ctx context.Context, userID uuid.UUID) ([]Reservation, error)
}
