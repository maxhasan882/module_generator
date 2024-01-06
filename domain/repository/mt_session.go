package repository

import (
	"context"
	"template/domain"
	"template/pkg/sson"
)

type CqExamRepo interface {
	Get(ctx context.Context, id string) (*domain.ModelTestResult, error)
	List(ctx context.Context, filter sson.D, skip int64, limit int64, sort ...interface{}) ([]*domain.ModelTestResult, error)
}
