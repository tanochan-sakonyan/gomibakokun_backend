package domain

import (
	domain "gomibakokun_backend/domain/trashcan"
	"testing"
)

func TestValidateTrashcanConfig(t *testing.T) {
	tests := []struct {
		name    string
		config  *domain.TrashcanConfig
		wantErr error
	}{
		{
			name: "Valid TrashcanConfig",
			config: &domain.TrashcanConfig{
				ID:              "123",
				Latitude:        35.6895,
				Longitude:       139.6917,
				Image:           "image_url",
				TrashType:       []string{"plastic"},
				NearestBuilding: "Tokyo Tower",
				Note:            "Near the park",
				SelectedButton:  "insideGate",
			},
			wantErr: nil,
		},
		{
			name: " Invalid TrashcanConfig - Empty ID",
			config: &domain.TrashcanConfig{
				ID:              "",
				Latitude:        35.6895,
				Longitude:       139.6917,
				Image:           "image_url",
				TrashType:       []string{"plastic"},
				NearestBuilding: "Tokyo Tower",
				Note:            "Near the park",
				SelectedButton:  "insideGate",
			},
			wantErr: domain.ErrInvalidInput,
		},
		{
			name: " Invalid TrashcanConfig - Invalid Latitude",
			config: &TrashcanConfig{
				ID:              "123",
				Latitude:        100.0,
				Longitude:       139.6917,
				Image:           "image_url",
				TrashType:       []string{"plastic"},
				NearestBuilding: "Tokyo Tower",
				Note:            "Near the park",
				SelectedButton:  "insideGate",
			},
			wantErr: ErrInvalidInput,
		},
		{
			name: " Invalid TrashcanConfig - Invalid Longitude",
			config: &TrashcanConfig{
				ID:              "123",
				Latitude:        35.6895,
				Longitude:       190.0,
				Image:           "image_url",
				TrashType:       []string{"plastic"},
				NearestBuilding: "Tokyo Tower",
				Note:            "Near the park",
				SelectedButton:  "insideGate",
			},
			wantErr: ErrInvalidInput,
		},
		{
			name: " Invalid TrashcanConfig - Invalid TrashType",
			config: &TrashcanConfig{
				ID:              "123",
				Latitude:        35.6895,
				Longitude:       160.0,
				Image:           "image_url",
				TrashType:       []string{"paper"},
				NearestBuilding: "Tokyo Tower",
				Note:            "Near the park",
				SelectedButton:  "insideGate",
			},
			wantErr: ErrInvalidInput,
		},
		{
			name: " Invalid TrashcanConfig - Invalid TrashType",
			config: &TrashcanConfig{
				ID:              "123",
				Latitude:        35.6895,
				Longitude:       -150.0,
				Image:           "image_url",
				TrashType:       []string{"paper"},
				NearestBuilding: "Tokyo Tower",
				Note:            "Near the park",
				SelectedButton:  "insideGate",
			},
			wantErr: ErrInvalidInput,
		},
		{
			name: " Invalid TrashcanConfig - Invalid SelectedButton",
			config: &TrashcanConfig{
				ID:              "123",
				Latitude:        35.6895,
				Longitude:       -150.0,
				Image:           "image_url",
				TrashType:       []string{"paper"},
				NearestBuilding: "Tokyo Tower",
				Note:            "Near the park",
				SelectedButton:  "recycle",
			},
			wantErr: ErrInvalidInput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTrashcanConfig(tt.config)
			if err != tt.wantErr {
				t.Errorf("ValidateTrashcanConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
