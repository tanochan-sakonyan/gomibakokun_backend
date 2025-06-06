package main

import (
	firestore "gomibakokun_backend/infrastructure"
	"gomibakokun_backend/infrastructure/persistence"
	"gomibakokun_backend/interfaces/handler"
	"gomibakokun_backend/usecase"

	"context"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	//Port番号を指定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//Firestoreの初期化
	ctx := context.Background()
	client, err := firestore.InitFirestoreClient(ctx, "gomibakokun")
	if err != nil {
		log.Fatalf("failed to initialize Firestore client: %v", err)
	}

	//依存性注入
	trashcanPersistence := persistence.NewTrashcanPersistence(client)
	trashcanUseCase := usecase.NewTrashcanUseCase(trashcanPersistence)
	trashcanHandler := handler.NewTrashcanHandler(trashcanUseCase)

	//Echoの初期化
	e := echo.New()
	//エンドポイントの設定
	e.POST("/trashcan", trashcanHandler.HandleTrashcanCreate)
	e.GET("/trashcan", trashcanHandler.HandleTrashcansInRange)
	e.DELETE("/trashcan", trashcanHandler.HandleTrashcanDelete)

	//サーバーの起動
	e.Logger.Fatal(e.Start(":" + port))
}
