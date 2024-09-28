package usecase

import (
	"bufio"
	"context"
	"fmt"
	"restApi/core/domain/dto"
	"strings"
)

// The following interfaces define the methods that the UseCaseMeliChallenge needs for its operations.
// These methods are implemented by different structs, such as AppSettings, MeliRepository, and SqlRepository,
// which serve as providers through the dependency injection pattern.

type FileConfig interface {
	IsAllowedSeparator(value string) (string, error)
}

type IReadItemFromRepository interface {
	GetItem(ctx context.Context, itemID string) (*dto.ItemDto, error)
}

type IWriteItemToRepository interface {
	Save(ctx context.Context, dto dto.ItemDto) error
}

// The UseCaseMeliChallenge is used to represent the main properties needed for the challenge.
type UseCaseMeliChallenge struct {
	config           FileConfig
	readItemFromRepo IReadItemFromRepository
	writeItemToRepo  IWriteItemToRepository
}

// NewUseCaseMeliChallenge is a constructor and a creational design pattern
// through it, dependencies needed by NewUseCaseMeliChallenge are injected for correct functionality.
func NewUseCaseMeliChallenge(config FileConfig, readItemFromRepo IReadItemFromRepository, writeItemToRepo IWriteItemToRepository) *UseCaseMeliChallenge {
	return &UseCaseMeliChallenge{
		config:           config,
		readItemFromRepo: readItemFromRepo,
		writeItemToRepo:  writeItemToRepo,
	}
}

// Execute is a method that runs over all items provided by the file
// for each row, it validates if it has an allowed separator and create goroutines to save in the repository.
func (uc *UseCaseMeliChallenge) Execute(ctx context.Context, arrayBytes []byte) error {

	scanner := bufio.NewScanner(strings.NewReader(string(arrayBytes)))

	scanner.Scan() // this moves to the next token,  Skip the header file

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		separator, err := uc.config.IsAllowedSeparator(line)
		// Validate separator
		if err == nil {

			itemID := strings.TrimSpace(strings.Replace(line, separator, "", 1))

			itemDto, err := uc.readItemFromRepo.GetItem(ctx, itemID)

			if err != nil {
				continue
			}

			if err := uc.writeItemToRepo.Save(ctx, *itemDto); err != nil {
				return fmt.Errorf("failed to save item: %w", err)
			}

		} else {
			return fmt.Errorf("invalid separator in line: %s", line)
		}
	}

	return nil
}
