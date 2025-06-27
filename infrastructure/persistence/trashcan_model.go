package persistence

import (
	domain "gomibakokun_backend/domain/trashcan"
)

type TrashcanModel struct {
	ID              string   `firestore:"id"`
	Latitude        float64  `firestore:"latitude"`
	Longitude       float64  `firestore:"longitude"`
	Image           string   `firestore:"image"`
	TrashType       []string `firestore:"trashType"`
	NearestBuilding string   `firestore:"nearestBuilding"`
	Note            string   `firestore:"note"`
	SelectedButton  string   `firestore:"selectedButton"`
}

func (t *TrashcanModel) toDomain() (*domain.Trashcan, error) {
	trashcan, err := domain.NewTrashcan(
		t.ID,
		t.Latitude,
		t.Longitude,
		t.Image,
		t.TrashType,
		t.NearestBuilding,
		t.Note,
		t.SelectedButton,
	)
	if err != nil {
		return nil, err
	}
	return trashcan, nil
}

func fromDomain(trashcan *domain.Trashcan) *TrashcanModel {
	return &TrashcanModel{
		ID:              trashcan.GetID(),
		Latitude:        trashcan.GetLatitude(),
		Longitude:       trashcan.GetLongitude(),
		Image:           trashcan.GetImage(),
		TrashType:       trashcan.GetTrashType(),
		NearestBuilding: trashcan.GetNearestBuilding(),
		Note:            trashcan.GetNote(),
		SelectedButton:  trashcan.GetSelectedButton(),
	}
}
