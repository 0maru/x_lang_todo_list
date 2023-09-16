package main

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/0maru/x_lang_todo_list/go_nextjs/ent"
	"github.com/0maru/x_lang_todo_list/go_nextjs/ent/todo"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
)

func Test_example(t *testing.T) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// タスクを作成する
	task1, err := client.Todo.Create().SetTitle("todo1").SetDescription("todo task1").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating todo: %v", err)
	}
	fmt.Println(task1)
	task2, err := client.Todo.Create().SetTitle("todo2").SetDescription("todo task2").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating todo: %v", err)
	}
	fmt.Println(task2)
	if err := task2.Update().SetParent(task1).Exec(ctx); err != nil {
		log.Fatalf("failed updating todo: %v", err)
	}

	// 全タスクを取得する
	items, err := client.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying todo: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n\n", t.ID, t.Title)
	}

	// 親タスクを持つタスクを取得する
	hasParentItems, err := client.Todo.Query().Where(todo.HasParent()).All(ctx)
	if err != nil {
		log.Fatalf("failed querying todo: %v", err)
	}
	for _, t := range hasParentItems {
		fmt.Printf("%d: %q\n\n", t.ID, t.Title)
	}

	// 親タスクを持たないかつ小タスクを持つタスクを取得する
	tasks, err := client.Todo.Query().Where(
		todo.Not(todo.HasParent()),
		todo.HasChildren(),
	).All(ctx)
	if err != nil {
		log.Fatalf("failed querying todo: %v", err)
	}
	for _, t := range tasks {
		fmt.Printf("%d: %q\n\n", t.ID, t.Title)
	}

	// 他のタスクを親に持つ子タスクから親タスクを取得する
	parent, err := client.Todo.Query().Where(todo.HasParent()).QueryParent().Only(ctx)
	if err != nil {
		log.Fatalf("failed querying todo: %v", err)
	}
	fmt.Printf("%d: %q\n\n", parent.ID, parent.Title)
}
