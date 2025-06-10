package domain

import (
	domain "gomibakokun_backend/domain/trashcan"
	"testing"
)

func TestValidateLongitude(t *testing.T) {
	tests := []struct {
		name      string
		longitude float64
		wantErr   error
	}{
		{
			name:      "Valid Longitude",
			longitude: 139.6917,
			wantErr:   nil,
		},
		{
			name:      "Invalid Longitude - Too Low",
			longitude: -181.0,
			wantErr:   domain.ErrInvalidInput,
		},
		{
			name:      "Invalid Longitude - Too High",
			longitude: 181.0,
			wantErr:   domain.ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := domain.ValidateLongitude(tt.longitude)
			if err != tt.wantErr {
				t.Errorf("ValidateLongitude() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
