package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restApi/core/domain/dto"
	"restApi/infra/config"
	"strconv"
	"sync"
)

type MeliRepository struct {
	client    http.Client
	EndPoints config.EndPoints
}

// Constructor: this is a design pattern that is creational
func NewMeliRepository(client http.Client, EndPoints config.EndPoints) *MeliRepository {
	return &MeliRepository{client: client, EndPoints: EndPoints}
}

// The GetItem function is responsible for retrieving items from several APIs. Additionally,
// it makes parallel calls to load the item DTO much faster
func (mr *MeliRepository) GetItem(ctx context.Context, itemID string) (*dto.ItemDto, error) {

	var item dto.ItemDto

	if err := mr.loadItem(&item, itemID); err != nil {
		return nil, fmt.Errorf("loading item: %v", err)
	}

	// Here, create a wait group for goroutines that are making API calls
	// to the orchestration service and populating the item DTO
	var wg sync.WaitGroup
	wg.Add(3)

	go func() (*dto.ItemDto, error) {
		defer wg.Done()
		if err := mr.loadSeller(&item); err != nil {
			return nil, fmt.Errorf("loading seller: %v", err)
		}

		return &item, nil
	}()

	go func() (*dto.ItemDto, error) {
		defer wg.Done()
		if err := mr.loadCategory(&item); err != nil {
			return nil, fmt.Errorf("loading category: %v", err)
		}

		return &item, nil
	}()

	go func() (*dto.ItemDto, error) {
		defer wg.Done()
		if err := mr.loadCurrency(&item); err != nil {
			return nil, fmt.Errorf("loading currency: %v", err)
		}
		return &item, nil
	}()

	wg.Wait()

	return &item, nil
}

// The following functions are auxiliary for loading data from several APIs.
func (mr *MeliRepository) loadItem(item *dto.ItemDto, itemID string) error {

	response, err := mr.client.Get(mr.EndPoints.ApiItems + itemID)
	if err != nil {
		return fmt.Errorf("getting item: %w", err)
	}

	if response == nil {
		return fmt.Errorf("Not Found item")
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		bytes, readErr := io.ReadAll(response.Body)
		if readErr != nil {
			return fmt.Errorf("reading Item: %w", readErr)
		}

		err = json.Unmarshal(bytes, &item)
		if err != nil {
			return fmt.Errorf("unmarshalling Item: %w", err)
		}

		return nil
	}

	return fmt.Errorf(response.Status)
}

func (mr *MeliRepository) loadSeller(item *dto.ItemDto) error {

	strSellerID := strconv.FormatUint(uint64(item.SellerID), 10)

	response, err := mr.client.Get(mr.EndPoints.ApiSeller + strSellerID)
	if err != nil {
		return fmt.Errorf("Error getting seller: %w", err)
	}
	if response == nil {
		return fmt.Errorf("Not Found seller")
	}
	defer response.Body.Close()

	bytes, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return fmt.Errorf("Error reading seller: %w", readErr)
	}

	err = json.Unmarshal(bytes, &item)
	if err != nil {
		return fmt.Errorf("Error Unmarshal seller: %w", readErr)
	}

	return nil

}

func (mr *MeliRepository) loadCategory(item *dto.ItemDto) error {

	response, err := mr.client.Get(mr.EndPoints.ApiCategory + item.CategoryID)
	if err != nil {
		return fmt.Errorf("Error getting category: %w", err)
	}
	if response == nil {
		return fmt.Errorf("Not Found category")
	}
	defer response.Body.Close()

	bytes, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return fmt.Errorf("Error reading category: %w", readErr)
	}
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		return fmt.Errorf("Error Unmarshal category: %w", readErr)
	}

	return nil
}

func (mr *MeliRepository) loadCurrency(item *dto.ItemDto) error {

	response, err := mr.client.Get(mr.EndPoints.ApiCurrency + item.CurrencyID)
	if err != nil {
		return fmt.Errorf("It was not connected to the currency API: %w", err)
	}
	if response == nil {
		return fmt.Errorf("Not Found currency")
	}
	defer response.Body.Close()

	bytes, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return fmt.Errorf("reading currency: %w", readErr)
	}

	err = json.Unmarshal(bytes, &item)
	if err != nil {
		return fmt.Errorf("value not found for currency: %w", readErr)
	}

	return nil
}
