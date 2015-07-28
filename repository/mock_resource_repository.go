package repository

import (
	"github.com/readit-tw/readit-api/model"
	"github.com/stretchr/testify/mock"
)

type MockResourceRepository struct {
	mock.Mock
}

func (m *MockResourceRepository) Create(r *model.Resource) (*model.Resource, error) {
	args := m.Called(r)
	return args.Get(0).(*model.Resource), args.Error(1)
}
