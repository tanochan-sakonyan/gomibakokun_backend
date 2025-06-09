package usecase

import (
	"context"
	"gomibakokun_backend/domain"
	"gomibakokun_backend/domain/repository"

	"math"

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
		if isInRange(latitude, longitude, trashcanLatitude, trashcanLongitude, radius) {
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

func isInRange(lat1, lon1, lat2, lon2, radiusKm float64) bool {
	// Haversine formula to calculate the distance between two points on the Earth
	const R = 6371 // Radius of the Earth in kilometers
	dLat := (lat2 - lat1) * (3.141592653589793 / 180)
	dLon := (lon2 - lon1) * (3.141592653589793 / 180)
	a := (math.Sin(dLat/2) * math.Sin(dLat/2)) + (math.Sin(lat1*(3.141592653589793/180)) * math.Sin(lat2*(3.141592653589793/180)) * math.Sin(dLon/2) * math.Sin(dLon/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c // Distance in kilometers

	return distance <= radiusKm
}
