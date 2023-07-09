package database

import "github.com/akifkadioglu/youtube-1/ent"

var client *ent.Client
var err error

type PostgreSQL struct{}
type MySQL struct{}
type SQLite struct{}

func Connection() {
	var database MySQL
	database.conn()
}

func DBManager() *ent.Client {
	return client
}
