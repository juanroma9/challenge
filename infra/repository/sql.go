package repository

import (
	"context"
	"database/sql"
	"restApi/core/domain/dto"
)

type SqlRepository struct {
	connection *sql.DB
}

func NewSQLRepository(connection *sql.DB) *SqlRepository {
	return &SqlRepository{connection: connection}
}

func (mr *SqlRepository) Save(ctx context.Context, dto dto.ItemDto) error {

	//rows, err := mr.connection.Query("insert into items (id) values(?)", itemID)
	rows, err := mr.connection.Query("INSERT INTO items (price, start_time, name, description, nickname) VALUES (?, ?, ?, ?, ?)", dto.Price, dto.DateCreated, dto.CategoryName, dto.CurrencyDescription, dto.SellerNickName)

	if err != nil {
		return err
	}

	rows.Close()

	return nil
}
