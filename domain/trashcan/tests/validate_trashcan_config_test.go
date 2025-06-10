package domain

import (
	domain "gomibakokun_backend/domain/trashcan"
	"testing"
)

func TestValidateTrashType(t *testing.T) {
	tests := []struct {
		name      string
		trashType []string
		wantErr   error
	}{
		{
			name:      "Valid Trash Type - burnable",
			trashType: []string{"burnable"},
			wantErr:   nil,
		},
		{
			name:      "Valid Trash Type - unburnable",
			trashType: []string{"unburnable"},
			wantErr:   nil,
		},
		{
			name:      "Valid Trash Type - pet_bottle",
			trashType: []string{"pet_bottle"},
			wantErr:   nil,
		},
		{
			name:      "Valid Trash Type - bottle",
			trashType: []string{"bottle"},
			wantErr:   nil,
		},
		{
			name:      "Valid Trash Type - can",
			trashType: []string{"can"},
			wantErr:   nil,
		},
		{
			name:      "Valid Trash Type - plastic",
			trashType: []string{"plastic"},
			wantErr:   nil,
		},
		{
			name:      "Valid Trash Type - other",
			trashType: []string{"other"},
			wantErr:   nil,
		},
		{
			name:      "Valid Trash Type - everything",
			trashType: []string{"everything"},
			wantErr:   nil,
		},
		{
			name:      "Valid Trash Type - ashtray",
			trashType: []string{"ashtray"},
			wantErr:   nil,
		},
		{
			name:      "Valid Trash Type - two types",
			trashType: []string{"burnable", "plastic"},
			wantErr:   nil,
		},
		{
			name:      "invalid Trash Type - unknown type",
			trashType: []string{"unknown"},
			wantErr:   domain.ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := domain.ValidateTrashType(tt.trashType)
			if err != tt.wantErr {
				t.Errorf("ValidateTrashType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
