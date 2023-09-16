package main

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/0maru/x_lang_todo_list/go_nextjs/ent"
	"log"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("failed opening connection to sqlite: ", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
