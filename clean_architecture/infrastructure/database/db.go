package database

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yona3/go-samples/clean_architecture/ent"
)

func NewDB() *ent.Client {
	db, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer db.Close()

	// オートマイグレーションツールを実行する
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return db
}
