package service

import (
	"context"
	"errors"
	"testing"

	"github.com/julioshinoda/port/domain"
	"github.com/julioshinoda/port/domain/mocks"
	"github.com/julioshinoda/port/entity"
	"github.com/stretchr/testify/mock"
)

func TestPortService_PortDomainService(t *testing.T) {
	type fields struct {
		portRepository domain.PortRepository
	}
	type args struct {
		ctx  context.Context
		port entity.Port
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Insert a registry",
			fields: fields{
				portRepository: &mocks.PortRepository{},
			},
			args: args{
				ctx: context.Background(),
				port: entity.Port{
					Name:    "Ajman",
					City:    "Ajman",
					Country: "United Arab Emirates",
					Alias:   []string{},
					Regions: []string{},
					Coordinates: []float64{
						55.5136433,
						25.4052165,
					},
					Province: "Ajman",
					Timezone: "Asia/Dubai",
					Unlocs:   []string{"AEAJM"},
					Code:     "52000",
				},
			},
			wantErr: false,
		},
		{
			name: "Error on upsert",
			fields: fields{
				portRepository: &mocks.PortRepository{},
			},
			args: args{
				ctx:  context.Background(),
				port: entity.Port{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PortService{
				portRepository: tt.fields.portRepository,
			}
			switch tt.name {
			case "Insert a registry":
				ps.portRepository.(*mocks.PortRepository).On("Upsert", mock.Anything, mock.Anything).Return(nil)
			case "Error on upsert":
				ps.portRepository.(*mocks.PortRepository).On("Upsert", mock.Anything, mock.Anything).Return(errors.New("error on database"))
			}
			if err := ps.PortDomainService(tt.args.ctx, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("PortService.PortDomainService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
