package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Todo.Type).Unique().From("children"),
	}
}
