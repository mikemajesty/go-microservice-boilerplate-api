package utils_tests

import (
	infra_repository "go-microservice-boilerplate-api/infra/repository"
	"go-microservice-boilerplate-api/utils"

	"github.com/stretchr/testify/mock"
)

type MockRepository[T utils.EntityAdapter, P utils.EntityIDAdapter] struct {
	mock.Mock
}

func (m *MockRepository[T, P]) Create(entity T, table string) (string, *utils.AppException) {
	args := m.Called(entity, table)
	if args.Get(1) != nil {
		return args.Get(0).(string), args.Get(1).(*utils.AppException)
	}
	return args.Get(0).(string), nil
}

func (m *MockRepository[T, P]) Delete(entity T, table string) *utils.AppException {
	args := m.Called(entity, table)
	return args.Get(0).(*utils.AppException)
}

func (m *MockRepository[T, P]) FindByID(input *infra_repository.FindOneInput[P], table string) (T, *utils.AppException) {
	args := m.Called(input, table)
	return args.Get(0).(T), args.Get(1).(*utils.AppException)
}

func (m *MockRepository[T, P]) List(table string) ([]T, *utils.AppException) {
	args := m.Called(table)
	return args.Get(0).([]T), args.Get(1).(*utils.AppException)
}

func (m *MockRepository[T, P]) Update(entity T, table string) (string, *utils.AppException) {
	args := m.Called(entity, table)
	return args.Get(0).(string), args.Get(1).(*utils.AppException)
}

func CreateBaseMock[T utils.EntityAdapter, P utils.EntityIDAdapter]() infra_repository.IRepository[T, P] {
	return &MockRepository[T, P]{}
}
