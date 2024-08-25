package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/realHoangHai/kratos/pkg/utils/entgo/mixin"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "role",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("name").
			NotEmpty().
			Optional(),

		field.String("guard_name").
			Comment("guard name").
			Optional(),

		field.Int32("description").
			Comment("description").
			Optional().
			Nillable(),
	}
}

// Mixin of the Role.
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.Timestamp{},
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
