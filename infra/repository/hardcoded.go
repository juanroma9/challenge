package repository

import (
	"context"
	"restApi/core/domain/dto"
)

// The following code was used to advance the development of the challenge
// while establishing the connection with the Meli APIs.

type HardcodedRepository struct{}

func NewHardcodedRepository() *HardcodedRepository {
	return &HardcodedRepository{}
}

func (hc *HardcodedRepository) GetItem(ctx context.Context, itemID string) (*dto.ItemDto, error) {
	return &dto.ItemDto{}, nil
}
