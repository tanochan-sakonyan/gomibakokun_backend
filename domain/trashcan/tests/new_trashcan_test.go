package domain

import (
	domain "gomibakokun_backend/domain/trashcan"
	"testing"
)

func TestNewTrashcan(t *testing.T) {
	tests := []struct {
		name            string
		id              string
		latitude        float64
		longitude       float64
		image           string
		trasyType       []string
		nearestBuilding string
		note            string
		selectedButton  string
		wantErr         error
	}{
		{
			name:            "Valid Trashcan",
			id:              "1",
			latitude:        35.6895,
			longitude:       139.6917,
			image:           "http://example.com/image.jpg",
			trasyType:       []string{"burnable", "unburnable"},
			nearestBuilding: "Test Building",
			note:            "Test Note",
			selectedButton:  "insideGate",
			wantErr:         nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trashcan, err := domain.NewTrashcan(
				tt.id,
				tt.latitude,
				tt.longitude,
				tt.image,
				tt.trasyType,
				tt.nearestBuilding,
				tt.note,
				tt.selectedButton,
			)
			if err != nil && err != tt.wantErr {
				t.Errorf("NewTrashcan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if trashcan.GetID() != tt.id {
					t.Errorf("Expected ID %s, got %s", tt.id, trashcan.GetID())
				}
			}
		})
	}
}
