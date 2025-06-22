package usecase

import (
	"context"
	"errors"
	"testing"

	usecase "gomibakokun_backend/usecase/trashcan"
	mock "gomibakokun_backend/usecase/trashcan/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteTrashcan_Success(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trashcanID := "123e4567-e89b-12d3-a456-426614174000"

	mockRepo := mock.NewMockTrashcanRepository(ctrl)
	mockRepo.EXPECT().DeleteTrashcan(ctx, trashcanID).Return(nil)

	trashcanUseCase := usecase.NewTrashcanUseCase(mockRepo)
	err := trashcanUseCase.DeleteTrashcan(ctx, trashcanID)
	require.NoError(t, err)
}

func TestDeleteTrashcan_Fail(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trashcanID := "123e4567-e89b-12d3-a456-426614174000"

	mockRepo := mock.NewMockTrashcanRepository(ctrl)
	mockRepo.EXPECT().DeleteTrashcan(ctx, trashcanID).Return(errors.New("not found"))

	trashcanUseCase := usecase.NewTrashcanUseCase(mockRepo)
	err := trashcanUseCase.DeleteTrashcan(ctx, trashcanID)

	// バリデーションエラーが発生することを検証
	require.Error(t, err)
}
