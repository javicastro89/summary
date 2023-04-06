package outbound

import (
	"context"
	"github.com/summary/internal/core/domain"
)

type DispatchMessageService interface {
	DispatchMessage(context.Context, domain.BalanceUser) error
}
