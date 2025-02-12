// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/realHoangHai/kratos/app/core/internal/data/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 1)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreateTime:  {Type: field.TypeInt64, Column: user.FieldCreateTime},
			user.FieldUpdateTime:  {Type: field.TypeInt64, Column: user.FieldUpdateTime},
			user.FieldDeleteTime:  {Type: field.TypeInt64, Column: user.FieldDeleteTime},
			user.FieldUsername:    {Type: field.TypeString, Column: user.FieldUsername},
			user.FieldPassword:    {Type: field.TypeString, Column: user.FieldPassword},
			user.FieldEmail:       {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldAvatar:      {Type: field.TypeString, Column: user.FieldAvatar},
			user.FieldDescription: {Type: field.TypeString, Column: user.FieldDescription},
			user.FieldAuthority:   {Type: field.TypeEnum, Column: user.FieldAuthority},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *UserFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreateTime applies the entql int64 predicate on the create_time field.
func (f *UserFilter) WhereCreateTime(p entql.Int64P) {
	f.Where(p.Field(user.FieldCreateTime))
}

// WhereUpdateTime applies the entql int64 predicate on the update_time field.
func (f *UserFilter) WhereUpdateTime(p entql.Int64P) {
	f.Where(p.Field(user.FieldUpdateTime))
}

// WhereDeleteTime applies the entql int64 predicate on the delete_time field.
func (f *UserFilter) WhereDeleteTime(p entql.Int64P) {
	f.Where(p.Field(user.FieldDeleteTime))
}

// WhereUsername applies the entql string predicate on the username field.
func (f *UserFilter) WhereUsername(p entql.StringP) {
	f.Where(p.Field(user.FieldUsername))
}

// WherePassword applies the entql string predicate on the password field.
func (f *UserFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(user.FieldPassword))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WhereAvatar applies the entql string predicate on the avatar field.
func (f *UserFilter) WhereAvatar(p entql.StringP) {
	f.Where(p.Field(user.FieldAvatar))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *UserFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(user.FieldDescription))
}

// WhereAuthority applies the entql string predicate on the authority field.
func (f *UserFilter) WhereAuthority(p entql.StringP) {
	f.Where(p.Field(user.FieldAuthority))
}
