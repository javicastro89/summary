package outbound

import (
	"context"
	"github.com/summary/internal/core/domain"
)

type RepositoryService interface {
	SaveItem(context.Context, domain.BalanceUser) error
}
