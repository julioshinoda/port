package domain

import (
	"context"

	"github.com/julioshinoda/port/entity"
)

type PortRepository interface {
	Upsert(ctx context.Context, port entity.Port) error
}
