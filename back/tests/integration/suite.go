package integration

import (
	"fmt"

	"github.com/alexinator1/sumb/back/internal/app"
	"github.com/alexinator1/sumb/back/tests"
	"github.com/alexinator1/sumb/back/tests/integration/helpers"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type IntegrationTestSuite struct {
	suite.Suite
	App *app.App
	db  *gorm.DB
	DbVerifier *helpers.DbVerifier
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.App = tests.NewApp()
	s.db = s.App.DB()
	s.DbVerifier = helpers.NewDbVerifier(s.db)
}

func (s *IntegrationTestSuite) ClearTables(tables ...string) {
	for _, table := range tables {
		s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
		s.db.Exec(fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART WITH 1", table))
	}
}
