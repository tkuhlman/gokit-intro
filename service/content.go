package service

import "context"

// Content defines the interface for our service.
type Content interface {
	Query(context.Context, string) (string, error)
}

type contentV1 struct{}

func (v1 *contentV1) Query(ctx context.Context, query string) (string, error) {
	return "Query Result", nil
}

// NewV1 returns a V1 implementation of the Content service.
func NewV1() Content {
	return &contentV1{}
}
