package usecase

import (
	"context"
	"errors"
	domain "gomibakokun_backend/domain/trashcan" // アサーションで利用
	"testing"
)

type mockTrashcanRepository struct {
	// 各メソッドの振る舞いをテストケースごとに設定するためのフィールド
	CreateTrashcanFunc func(ctx context.Context, trashcan *domain.Trashcan) error
	GetAllTrashcanFunc func(ctx context.Context) ([]*domain.Trashcan, error)
	DeleteTrashcanFunc func(ctx context.Context, id string) error
}

// repository.TrashcanRepositoryインターフェースのメソッドを実装します。
// 実際にはフィールドに設定された関数を呼び出すだけです。

func (m *mockTrashcanRepository) CreateTrashcan(ctx context.Context, trashcan *domain.Trashcan) error {
	if m.CreateTrashcanFunc != nil {
		return m.CreateTrashcanFunc(ctx, trashcan)
	}
	// デフォルトの振る舞い（何もしない、エラーなし）
	return nil
}

func (m *mockTrashcanRepository) GetAllTrashcan(ctx context.Context) ([]*domain.Trashcan, error) {
	if m.GetAllTrashcanFunc != nil {
		return m.GetAllTrashcanFunc(ctx)
	}
	return nil, nil
}

func TestCreateTrashcan(t *testing.T) {
	validTrashcanConfig := &domain.TrashcanConfig{
		ID:              "123",
		Latitude:        35.6895,
		Longitude:       139.6917,
		Image:           "image_url",
		TrashType:       []string{"burnable"},
		NearestBuilding: "Tokyo Tower",
		Note:            "Near the park",
		SelectedButton:  "insideGate",
	}

	invalidTrashcanConfig := &domain.TrashcanConfig{
		ID:              "",
		Latitude:        35.6895,
		Longitude:       139.6917,
		Image:           "image_url",
		TrashType:       []string{"burnable"},
		NearestBuilding: "Tokyo Tower",
		Note:            "Near the park",
		SelectedButton:  "insideGate",
	}

	tests := []struct {
		name           string
		trashcanConfig *domain.TrashcanConfig
		setupMock      func(mock *mockTrashcanRepository)
		expectedError  error
	}{
		{
			name:           "正常系：ゴミ箱の作成が成功する",
			trashcanConfig: validTrashcanConfig,
			setupMock: func(mock *mockTrashcanRepository) {
				mock.CreateTrashcanFunc = func(ctx context.Context, trashcan *domain.Trashcan) error {
					return nil // 成功
				}
			},
			expectedError: nil,
		},
		{
			name:           "異常系：リポジトリ層での失敗",
			trashcanConfig: validTrashcanConfig,
			setupMock: func(mock *mockTrashcanRepository) {
				mock.CreateTrashcanFunc = func(ctx context.Context, trashcan *domain.Trashcan) error {
					return errors.New("fail") // リポジトリ層でのエラー
				}
			},
			expectedError: errors.New("fail"),
		},
		{
			name:           "異常系：無効な入力",
			trashcanConfig: invalidTrashcanConfig,
			setupMock:      func(mock *mockTrashcanRepository) {},
			expectedError:  domain.ErrInvalidInput,
		},
	}

}
