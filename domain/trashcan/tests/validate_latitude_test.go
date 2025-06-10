package domain

import (
	domain "gomibakokun_backend/domain/trashcan"
	"testing"
)

func TestValidateLatitude(t *testing.T) {
	tests := []struct {
		name     string
		latitude float64
		wantErr  error
	}{
		{
			name:     "Valid Latitude",
			latitude: 35.6895,
			wantErr:  nil,
		},
		{
			name:     "Invalid Latitude - Too Low",
			latitude: -91.0,
			wantErr:  domain.ErrInvalidInput,
		},
		{
			name:     "Invalid Latitude - Too High",
			latitude: 91.0,
			wantErr:  domain.ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := domain.ValidateLatitude(tt.latitude)
			if err != tt.wantErr {
				t.Errorf("ValidateLatitude() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
