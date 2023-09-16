package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MaxLen(255),
		field.Text("description"),
		field.Enum("status").
			Values("todo", "in_progress", "done"),
		field.Time("created_at").
			Default(time.Now),
	}
}
