package persistence

import (
	"context"
	"gomibakokun_backend/domain/repository"
	domain "gomibakokun_backend/domain/trashcan"
	"log"

	"cloud.google.com/go/firestore"
)

type trashcanPersistence struct {
	client *firestore.Client
}

func NewTrashcanPersistence(client *firestore.Client) repository.TrashcanRepository {
	return &trashcanPersistence{client: client}
}

func (tp trashcanPersistence) CreateTrashcan(ctx context.Context, trashcan *domain.Trashcan) error {
	_, err := tp.client.Collection("trashcans").Doc(trashcan.GetID()).Set(ctx, trashcan)
	if err != nil {
		log.Printf("An error has occurred to create trashcan: %s", err)
	}

	return err
}

func (tp trashcanPersistence) GetAllTrashcan(ctx context.Context) ([]*domain.Trashcan, error) {
	iter := tp.client.Collection("trashcans").Documents(ctx)
	var trashcans []*domain.Trashcan

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var trashcan domain.Trashcan
		if err := doc.DataTo(&trashcan); err != nil {
			log.Printf("An error has occurred to get all trashcans: %s", err)
			return nil, err
		}
		trashcans = append(trashcans, &trashcan)
	}

	return trashcans, nil
}

func (tp trashcanPersistence) DeleteTrashcan(ctx context.Context, id string) error {
	_, err := tp.client.Collection("trashcans").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("An error has occurred to delete trashcan: %s", err)
	}
	return err
}
