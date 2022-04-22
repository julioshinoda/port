package service

import (
	"context"

	"github.com/julioshinoda/port/domain"
	"github.com/julioshinoda/port/domain/postgres"
	"github.com/julioshinoda/port/entity"
)

type PortService struct {
	portRepository domain.PortRepository
}

//NewPortService is a factory to create a new Port service
func NewPortService(ctx context.Context, dbURL string) (*PortService, error) {
	conn, err := postgres.NewPostgresRepository(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	ps := &PortService{
		portRepository: conn,
	}

	return ps, nil
}

func (ps *PortService) PortDomainService(ctx context.Context, port entity.Port) error {
	return ps.portRepository.Upsert(ctx, port)
}
