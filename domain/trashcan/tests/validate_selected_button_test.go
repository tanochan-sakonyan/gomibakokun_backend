package domain

import (
	domain "gomibakokun_backend/domain/trashcan"
	"testing"
)

func TestValidateSelectedButton(t *testing.T) {
	tests := []struct {
		name           string
		selectedButton string
		wantErr        error
	}{
		{
			name:           "Valid SelectedButton - insideGate",
			selectedButton: "insideGate",
			wantErr:        nil,
		},
		{
			name:           "Valid SelectedButton - outside",
			selectedButton: "outside",
			wantErr:        nil,
		},
		{
			name:           "Valid SelectedButton - insideBuilding",
			selectedButton: "insideBuilding",
			wantErr:        nil,
		},
		{
			name:           "Invalid SelectedButton - empty string",
			selectedButton: "",
			wantErr:        domain.ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := domain.ValidateSelectedButton(tt.selectedButton)
			if err != tt.wantErr {
				t.Errorf("ValidateSelectedButton() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
