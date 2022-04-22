package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julioshinoda/port/domain"
	"github.com/julioshinoda/port/entity"
	"github.com/julioshinoda/port/infrastructure/postgres"
	"github.com/lib/pq"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(ctx context.Context, connectionString string) (domain.PortRepository, error) {
	conn, err := postgres.CreateConnection(connectionString)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{
		db: conn,
	}, nil
}

func (pr *PostgresRepository) Upsert(ctx context.Context, port entity.Port) error {
	_, err := pr.db.Exec(
		context.Background(),
		upsertQuery,
		port.ID,
		port.Name,
		port.City,
		port.Country,
		pq.Array(port.Alias),
		pq.Array(port.Regions),
		pq.Array(port.Coordinates),
		port.Province,
		port.Timezone,
		pq.Array(port.Unlocs),
		port.Code,
	)
	return err
}

var upsertQuery = `INSERT INTO public.port (
						id,
						"name",
						city,
						country,
						alias,
						regions,
						coordinates,
						province,
						timezone,
						unlocs,
						code
					) VALUES (
						$1,
						$2,
						$3,
						$4,
						$5,
						$6,
						$7,
						$8,
						$9,
						$10,
						$11
						) ON CONFLICT (id) DO 
					UPDATE SET  
					   "name"      = $2,
					   city        = $3,
					   country     = $4,
					   alias       = $5,
					   regions     = $6,
					   coordinates = $7,
					   province    = $8,
					   timezone    = $9,
					   unlocs      = $10,
					   code        = $11
					   `
