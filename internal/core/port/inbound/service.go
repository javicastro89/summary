package inbound

import (
	"context"
	"github.com/summary/internal/core/domain"
)

type ProcessSummaryService interface {
	ProcessSummary(context.Context, []domain.User) error
}
