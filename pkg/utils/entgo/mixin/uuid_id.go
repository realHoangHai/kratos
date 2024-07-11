package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

var _ ent.Mixin = (*UUIDId)(nil)

type UUIDId struct{ mixin.Schema }

func (UUIDId) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("id").
			Default(uuid.New).
			Unique().
			Immutable().
			StructTag(`json:"id,omitempty"`),
	}
}

// Indexes of the UuidId.
func (UUIDId) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
