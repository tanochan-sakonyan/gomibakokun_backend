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
	trashcanModel := fromDomain(trashcan)
	if trashcanModel == nil {
		log.Printf("An error has occurred to convert trashcan to model: %s", trashcan.GetID())
		return domain.ErrInvalidInput
	}

	_, err := tp.client.Collection("trashcans").Doc(trashcan.GetID()).Set(ctx, trashcanModel)
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
		var trashcanModel TrashcanModel
		if err := doc.DataTo(&trashcanModel); err != nil {
			log.Printf("An error has occurred to convert document to trashcan model: %s", err)
		}
		log.Printf("Document ID: %s", doc.Ref.ID)
		trashcan, err := trashcanModel.toDomain()
		if err != nil {
			log.Printf("An error has occurred to convert trashcan model to domain: %s", err)
		}

		trashcans = append(trashcans, trashcan)
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
