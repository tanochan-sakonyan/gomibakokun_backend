package persistence

import (
	"testing"
)

func TestTrashcanModelToDomain(t *testing.T) {
	trashcanModel := &TrashcanModel{
		ID:              "test-id",
		Latitude:        35.6895,
		Longitude:       139.6917,
		Image:           "http://example.com/image.jpg",
		TrashType:       []string{"burnable", "unburnable"},
		NearestBuilding: "Test Building",
		Note:            "Test Note",
		SelectedButton:  "burnable",
	}

	trashcanDomain, err := trashcanModel.toDomain()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if trashcanDomain.GetID() != trashcanModel.ID {
		t.Errorf("Expected ID %s, got %s", trashcanModel.ID, trashcanDomain.GetID())
	}
}
