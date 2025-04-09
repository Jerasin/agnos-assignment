package mocks

import (
	"agnos-assignment/app/repository"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockPatientRepository struct {
	mock.Mock
}

func (m *MockPatientRepository) GetBaseRepo() repository.BaseRepositoryInterface {
	fmt.Println("GetBaseRepo")
	args := m.Called()
	return args.Get(0).(repository.BaseRepositoryInterface)
}
