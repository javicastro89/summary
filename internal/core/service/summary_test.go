package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/summary/internal/core/domain"
	"testing"
)

var sliceUser = []domain.User{
	{
		Id:          "som@thing.com",
		Date:        "7/15",
		Transaction: -3.14,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: +8.14,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: +115.25,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: -865.34,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: +1115.25,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: +156.14,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: -550.15,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: +56,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: -865.34,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: +56,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: +15,
	},
	{
		Id:          "som@thing.com",
		Date:        "9/15",
		Transaction: +1115.25,
	},
}

func TestService_ProcessSummary(t *testing.T) {
	mkDispatch := new(mockDispatch)
	mkDispatch.On("DispatchMessage", mock.AnythingOfType("domain.BalanceUser")).Return(nil)

	mkRepository := new(mockRepository)
	mkRepository.On("SaveItem", mock.AnythingOfType("domain.BalanceUser")).Return(nil)

	srv := NewService(mkDispatch, mkRepository)

	err := srv.ProcessSummary(context.Background(), sliceUser)
	assert.NoError(t, err)
}

// ------- Mocks ---------
type mockDispatch struct {
	mock.Mock
}

func (m *mockDispatch) DispatchMessage(_ context.Context, balanceUser domain.BalanceUser) error {
	args := m.Called(balanceUser)
	return args.Error(0)
}

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) SaveItem(_ context.Context, balanceUser domain.BalanceUser) error {
	args := m.Called(balanceUser)
	return args.Error(0)
}
