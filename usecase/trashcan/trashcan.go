package usecase

import (
	"context"
	"gomibakokun_backend/domain/repository"
	domain "gomibakokun_backend/domain/trashcan"

	"github.com/google/uuid"
)

type TrashcanUseCase interface {
	CreateTrashcan(ctx context.Context, trashcanConfig *domain.TrashcanConfig) error
	GetTrashcansInRange(ctx context.Context, latitude float64, longitude float64, radius float64) ([]*domain.Trashcan, error)
	DeleteTrashcan(ctx context.Context, id string) error
}

type trashcanUseCase struct {
	trashcanRepository repository.TrashcanRepository
}

func NewTrashcanUseCase(tr repository.TrashcanRepository) TrashcanUseCase {
	return &trashcanUseCase{
		trashcanRepository: tr,
	}
}

func (tu *trashcanUseCase) CreateTrashcan(ctx context.Context, trashcanConfig *domain.TrashcanConfig) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	trashcanConfig.ID = id.String()

	//ゴミ箱ドメインの作成
	trashcan, err := domain.NewTrashcan(trashcanConfig)
	if err != nil {
		return domain.ErrInvalidInput
	}

	err = tu.trashcanRepository.CreateTrashcan(ctx, trashcan)

	return err
}

func (tu *trashcanUseCase) GetTrashcansInRange(ctx context.Context, latitude float64, longitude float64, radius float64) ([]*domain.Trashcan, error) {
	trashcans, err := tu.trashcanRepository.GetAllTrashcan(ctx)
	if err != nil {
		return nil, err
	}

	var trashcansInRange []*domain.Trashcan
	for _, trashcan := range trashcans {
		trashcanLatitude, trashcanLongitude := trashcan.GetLatitudeAndLongitude()
		if IsInRange(latitude, longitude, trashcanLatitude, trashcanLongitude, radius) {
			trashcansInRange = append(trashcansInRange, trashcan)
		}
	}

	return trashcansInRange, nil
}

func (tu *trashcanUseCase) DeleteTrashcan(ctx context.Context, id string) error {
	err := tu.trashcanRepository.DeleteTrashcan(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
