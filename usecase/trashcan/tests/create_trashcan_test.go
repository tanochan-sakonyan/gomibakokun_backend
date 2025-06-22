package usecase

import (
	"context"
	"testing"

	domain "gomibakokun_backend/domain/trashcan"
	usecase "gomibakokun_backend/usecase/trashcan"
	mock "gomibakokun_backend/usecase/trashcan/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateTrashcan_Success(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trashcanConfig := &domain.TrashcanConfig{
		Latitude:        350.6895,
		Longitude:       139.6917,
		Image:           "image_url",
		TrashType:       []string{"burnable"},
		NearestBuilding: "Tokyo Tower",
		Note:            "Near the park",
		SelectedButton:  "insideGate",
	}

	mockRepo := mock.NewMockTrashcanRepository(ctrl)
	mockRepo.EXPECT().CreateTrashcan(ctx, gomock.Any()).Return(nil)
	trashcanUseCase := usecase.NewTrashcanUseCase(mockRepo)
	err := trashcanUseCase.CreateTrashcan(ctx, trashcanConfig)
	require.NoError(t, err)
}

func TestCreateTrashcan_Fail(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trashcanConfig := &domain.TrashcanConfig{
		Latitude:  350.6895, // 不正な値
		Longitude: 139.6917,
		// ...
	}

	mockRepo := mock.NewMockTrashcanRepository(ctrl)
	// mockRepo.EXPECT() の記述は不要

	trashcanUseCase := usecase.NewTrashcanUseCase(mockRepo)
	err := trashcanUseCase.CreateTrashcan(ctx, trashcanConfig)

	// バリデーションエラーが発生することを検証
	require.ErrorIs(t, err, domain.ErrInvalidInput)
}
