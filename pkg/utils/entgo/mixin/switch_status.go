package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type SwitchStatus struct {
	mixin.Schema
}

func (SwitchStatus) Fields() []ent.Field {
	return []ent.Field{
		/**
		Under PostgreSQL, you also need to create a Type for this, otherwise it cannot be used.

		DROP TYPE IF EXISTS switch_status CASCADE;
		CREATE TYPE switch_status AS ENUM (
		    'OFF',
		    'ON'
		    );
		*/
		field.Enum("status").
			Comment("Status").
			Optional().
			Nillable().
			//SchemaType(map[string]string{
			//	dialect.MySQL:    "switch_status",
			//	dialect.Postgres: "switch_status",
			//}).
			Default("ON").
			Values(
				"OFF",
				"ON",
			),
	}
}
