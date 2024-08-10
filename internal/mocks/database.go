package mocks

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
	"marketplace/internal/models"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Prepare(query string) (models.Stmt, error) {
	args := m.Called(query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(models.Stmt), args.Error(1)
}

type MockStmt struct {
	mock.Mock
}

func (m *MockStmt) Exec(args ...interface{}) (sql.Result, error) {
	arguments := m.Called(args...)
	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}
	return arguments.Get(0).(sql.Result), arguments.Error(1)
}

func (m *MockStmt) Close() error {
	return nil
}
