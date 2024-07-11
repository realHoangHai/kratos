package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	mixin2 "github.com/realHoangHai/kratos/pkg/utils/entgo/mixin"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "user",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Comment("username").
			Unique().
			MaxLen(50).
			NotEmpty().
			Immutable().
			Optional().
			Nillable().
			Match(regexp.MustCompile("^[a-zA-Z0-9]{4,16}$")),

		field.String("password").
			Comment("password").
			MaxLen(255).
			Optional().
			Nillable().
			NotEmpty(),

		field.String("email").
			Comment("email").
			MaxLen(127).
			Optional().
			Nillable(),

		field.String("avatar").
			Comment("avatar").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.String("description").
			Comment("description").
			MaxLen(1023).
			Optional().
			Nillable(),

		field.Enum("authority").
			Comment("authorize").
			Optional().
			Nillable().
			//SchemaType(map[string]string{
			//	dialect.MySQL:    "authority",
			//	dialect.Postgres: "authority",
			//}).
			Values(
				"SYS_ADMIN",
				"CUSTOMER_USER",
				"GUEST_USER",
				"REFRESH_TOKEN",
			).
			Default("CUSTOMER_USER"),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin2.AutoIncrementId{},
		mixin2.Timestamp{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "username").Unique(),
	}
}
