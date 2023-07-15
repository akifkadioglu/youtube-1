package database

import (
	"context"
	"log"

	"github.com/akifkadioglu/youtube-1/ent"
	"github.com/akifkadioglu/youtube-1/env"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

)

func (MySQL) conn() {
	client, err = ent.Open("mysql", env.Getenv(env.DB_USERNAME)+":"+env.Getenv(env.DB_PASSWORD)+"@tcp("+env.Getenv(env.DB_HOST)+":"+env.Getenv(env.DB_PORT)+")/"+env.Getenv(env.DB_DATABASE)+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("connected to database")
}

func (SQLite) conn() {
	client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}