package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*CreateTimestamp)(nil)

type CreateTimestamp struct{ mixin.Schema }

func (CreateTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// create time - milliseconds
		field.Int64("create_time").
			Comment("create time").
			Immutable().
			Optional().
			Nillable().
			DefaultFunc(time.Now().UnixMilli),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*UpdateTimestamp)(nil)

type UpdateTimestamp struct{ mixin.Schema }

func (UpdateTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// update time, milliseconds
		// It should be noted that if the program does not update automatically, then this field will not be updated unless a trigger is updated in the database.
		field.Int64("update_time").
			Comment("update time").
			Optional().
			Nillable().
			UpdateDefault(time.Now().UnixMilli),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*DeleteTimestamp)(nil)

type DeleteTimestamp struct{ mixin.Schema }

func (DeleteTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// delete time, milliseconds
		field.Int64("delete_time").
			Comment("delete time").
			Optional().
			Nillable(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*Timestamp)(nil)

type Timestamp struct{ mixin.Schema }

func (Timestamp) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, CreateTimestamp{}.Fields()...)
	fields = append(fields, UpdateTimestamp{}.Fields()...)
	fields = append(fields, DeleteTimestamp{}.Fields()...)
	return fields
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*CreatedAtTimestamp)(nil)

type CreatedAtTimestamp struct{ mixin.Schema }

func (CreatedAtTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// create time - milliseconds
		field.Int64("created_at").
			Comment("create time").
			Immutable().
			Optional().
			Nillable().
			DefaultFunc(time.Now().UnixMilli),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*UpdatedAtTimestamp)(nil)

type UpdatedAtTimestamp struct{ mixin.Schema }

func (UpdatedAtTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// update time, milliseconds
		// It should be noted that if the program does not update automatically, then this field will not be updated unless a trigger is updated in the database.
		field.Int64("updated_at").
			Comment("update time").
			Optional().
			Nillable().
			UpdateDefault(time.Now().UnixMilli),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*DeletedAtTimestamp)(nil)

type DeletedAtTimestamp struct{ mixin.Schema }

func (DeletedAtTimestamp) Fields() []ent.Field {
	return []ent.Field{
		// delete time, milliseconds
		field.Int64("deleted_at").
			Comment("delete time").
			Optional().
			Nillable(),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ ent.Mixin = (*TimestampAt)(nil)

type TimestampAt struct{ mixin.Schema }

func (TimestampAt) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, CreatedAtTimestamp{}.Fields()...)
	fields = append(fields, UpdatedAtTimestamp{}.Fields()...)
	fields = append(fields, DeletedAtTimestamp{}.Fields()...)
	return fields
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
