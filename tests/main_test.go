package tests

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"houseService/internal/app"
	"houseService/pkg/auth"
	"log/slog"
	"os"
	"testing"
	"time"
)

type APITestSuite struct {
	suite.Suite

	db *sqlx.DB

	tokenManager *auth.Manager

	serviceProvider *app.ServiceProvider
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) SetupSuite() {
	dataSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		"test", "test", "localhost", "5430", "homeServiceTest")

	if db, err := sqlx.Connect("postgres", dataSource); err != nil {
		s.FailNow("Failed to connect to postgres", err)
	} else {
		s.db = db
	}

	s.initDeps()

	if err := s.initDB(); err != nil {
		s.FailNow("Failed to create and populate DB", err)
	}
}

func (s *APITestSuite) TearDownSuite() {
	s.db.Close()
}

func (s *APITestSuite) initDeps() {
	// Init domain deps
	d, _ := time.ParseDuration("2h")
	if tokenManager, err := auth.NewManager("signing_key", d); err != nil {
		s.FailNow("Failed to initialize token manager", err)
	} else {
		s.tokenManager = tokenManager
	}

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	s.serviceProvider = app.NewServiceProvider(log, s.db, s.tokenManager)

	s.serviceProvider.RegisterControllers()

}

func (s *APITestSuite) initDB() error {
	createHouses := `	CREATE TABLE IF NOT EXISTS houses(
				house_number SERIAL PRIMARY KEY,
				address VARCHAR NOT NULL,
				year_construction INTEGER NOT NULL,
				developer VARCHAR,
				created_at TIMESTAMP NOT NULL,
				last_update TIMESTAMP NOT NULL)`

	_, err := s.db.Exec(createHouses)
	if err != nil {
		s.FailNow("Failed to create houses table", err)
	}

	insertHouse := `INSERT INTO houses VALUES (1, 'address', 2000, '', '2016-06-22 19:10:25-07' , '2016-06-22 19:10:25-07');`

	_, err = s.db.Exec(insertHouse)
	if err != nil {
		s.FailNow("Failed to insert house", err)
	}

	createFlats := `	CREATE TABLE IF NOT EXISTS flats(
						flat_number SERIAL NOT NULL,
						house_number INTEGER REFERENCES houses ON DELETE CASCADE,
						price INTEGER NOT NULL,
						number_of_rooms INTEGER NOT NULL,
						status VARCHAR NOT NULL,
						PRIMARY KEY (flat_number, house_number))`

	_, err = s.db.Exec(createFlats)
	if err != nil {
		s.FailNow("Failed to create flats table", err)
	}

	insertFlats := `INSERT INTO flats VALUES (11, 1, 20, 1 , 'created'), (22, 1, 30, 2 , 'approved')`

	_, err = s.db.Exec(insertFlats)
	if err != nil {
		s.FailNow("Failed to insert flats", err)
	}

	return nil
}
