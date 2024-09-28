package repository

import (
	"context"
	"restApi/core/domain/dto"
)

// The following code was used to advance the development of the challenge
// while establishg the connection with the Meli APIs.
type MemoryRepository struct {
	items map[string]string
}

func NewMemoryRepository(items map[string]string) *MemoryRepository {
	return &MemoryRepository{
		items: loadItem(),
	}
}

func (mr *MemoryRepository) GetItem(ctx context.Context, itemID string) (*dto.ItemDto, error) {
	return &dto.ItemDto{}, nil
}

func loadItem() map[string]string {
	return map[string]string{
		"MLA750925229": " Item de MLA con ID 750925229",
		"MLA845041373": " Item de MLA con ID 845041373",
		"MLA693105237": " Item de MLA con ID 693105237",
	}
}
