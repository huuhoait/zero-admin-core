// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/menu"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/role"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RoleUpdate) SetUpdatedAt(t time.Time) *RoleUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetStatus sets the "status" field.
func (ru *RoleUpdate) SetStatus(u uint8) *RoleUpdate {
	ru.mutation.ResetStatus()
	ru.mutation.SetStatus(u)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableStatus(u *uint8) *RoleUpdate {
	if u != nil {
		ru.SetStatus(*u)
	}
	return ru
}

// AddStatus adds u to the "status" field.
func (ru *RoleUpdate) AddStatus(u int8) *RoleUpdate {
	ru.mutation.AddStatus(u)
	return ru
}

// ClearStatus clears the value of the "status" field.
func (ru *RoleUpdate) ClearStatus() *RoleUpdate {
	ru.mutation.ClearStatus()
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetCode sets the "code" field.
func (ru *RoleUpdate) SetCode(s string) *RoleUpdate {
	ru.mutation.SetCode(s)
	return ru
}

// SetDefaultRouter sets the "default_router" field.
func (ru *RoleUpdate) SetDefaultRouter(s string) *RoleUpdate {
	ru.mutation.SetDefaultRouter(s)
	return ru
}

// SetNillableDefaultRouter sets the "default_router" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableDefaultRouter(s *string) *RoleUpdate {
	if s != nil {
		ru.SetDefaultRouter(*s)
	}
	return ru
}

// SetRemark sets the "remark" field.
func (ru *RoleUpdate) SetRemark(s string) *RoleUpdate {
	ru.mutation.SetRemark(s)
	return ru
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableRemark(s *string) *RoleUpdate {
	if s != nil {
		ru.SetRemark(*s)
	}
	return ru
}

// SetSort sets the "sort" field.
func (ru *RoleUpdate) SetSort(u uint32) *RoleUpdate {
	ru.mutation.ResetSort()
	ru.mutation.SetSort(u)
	return ru
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableSort(u *uint32) *RoleUpdate {
	if u != nil {
		ru.SetSort(*u)
	}
	return ru
}

// AddSort adds u to the "sort" field.
func (ru *RoleUpdate) AddSort(u int32) *RoleUpdate {
	ru.mutation.AddSort(u)
	return ru
}

// AddMenuIDs adds the "menus" edge to the Menu entity by IDs.
func (ru *RoleUpdate) AddMenuIDs(ids ...uint64) *RoleUpdate {
	ru.mutation.AddMenuIDs(ids...)
	return ru
}

// AddMenus adds the "menus" edges to the Menu entity.
func (ru *RoleUpdate) AddMenus(m ...*Menu) *RoleUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ru.AddMenuIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ru *RoleUpdate) AddUserIDs(ids ...uuid.UUID) *RoleUpdate {
	ru.mutation.AddUserIDs(ids...)
	return ru
}

// AddUsers adds the "users" edges to the User entity.
func (ru *RoleUpdate) AddUsers(u ...*User) *RoleUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.AddUserIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// ClearMenus clears all "menus" edges to the Menu entity.
func (ru *RoleUpdate) ClearMenus() *RoleUpdate {
	ru.mutation.ClearMenus()
	return ru
}

// RemoveMenuIDs removes the "menus" edge to Menu entities by IDs.
func (ru *RoleUpdate) RemoveMenuIDs(ids ...uint64) *RoleUpdate {
	ru.mutation.RemoveMenuIDs(ids...)
	return ru
}

// RemoveMenus removes "menus" edges to Menu entities.
func (ru *RoleUpdate) RemoveMenus(m ...*Menu) *RoleUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ru.RemoveMenuIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (ru *RoleUpdate) ClearUsers() *RoleUpdate {
	ru.mutation.ClearUsers()
	return ru
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (ru *RoleUpdate) RemoveUserIDs(ids ...uuid.UUID) *RoleUpdate {
	ru.mutation.RemoveUserIDs(ids...)
	return ru
}

// RemoveUsers removes "users" edges to User entities.
func (ru *RoleUpdate) RemoveUsers(u ...*User) *RoleUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks[int, RoleMutation](ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoleUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := role.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint64))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := ru.mutation.AddedStatus(); ok {
		_spec.AddField(role.FieldStatus, field.TypeUint8, value)
	}
	if ru.mutation.StatusCleared() {
		_spec.ClearField(role.FieldStatus, field.TypeUint8)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if value, ok := ru.mutation.DefaultRouter(); ok {
		_spec.SetField(role.FieldDefaultRouter, field.TypeString, value)
	}
	if value, ok := ru.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
	}
	if value, ok := ru.mutation.Sort(); ok {
		_spec.SetField(role.FieldSort, field.TypeUint32, value)
	}
	if value, ok := ru.mutation.AddedSort(); ok {
		_spec.AddField(role.FieldSort, field.TypeUint32, value)
	}
	if ru.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: role.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedMenusIDs(); len(nodes) > 0 && !ru.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: role.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: role.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !ru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RoleUpdateOne) SetUpdatedAt(t time.Time) *RoleUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RoleUpdateOne) SetStatus(u uint8) *RoleUpdateOne {
	ruo.mutation.ResetStatus()
	ruo.mutation.SetStatus(u)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableStatus(u *uint8) *RoleUpdateOne {
	if u != nil {
		ruo.SetStatus(*u)
	}
	return ruo
}

// AddStatus adds u to the "status" field.
func (ruo *RoleUpdateOne) AddStatus(u int8) *RoleUpdateOne {
	ruo.mutation.AddStatus(u)
	return ruo
}

// ClearStatus clears the value of the "status" field.
func (ruo *RoleUpdateOne) ClearStatus() *RoleUpdateOne {
	ruo.mutation.ClearStatus()
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetCode sets the "code" field.
func (ruo *RoleUpdateOne) SetCode(s string) *RoleUpdateOne {
	ruo.mutation.SetCode(s)
	return ruo
}

// SetDefaultRouter sets the "default_router" field.
func (ruo *RoleUpdateOne) SetDefaultRouter(s string) *RoleUpdateOne {
	ruo.mutation.SetDefaultRouter(s)
	return ruo
}

// SetNillableDefaultRouter sets the "default_router" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableDefaultRouter(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetDefaultRouter(*s)
	}
	return ruo
}

// SetRemark sets the "remark" field.
func (ruo *RoleUpdateOne) SetRemark(s string) *RoleUpdateOne {
	ruo.mutation.SetRemark(s)
	return ruo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableRemark(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetRemark(*s)
	}
	return ruo
}

// SetSort sets the "sort" field.
func (ruo *RoleUpdateOne) SetSort(u uint32) *RoleUpdateOne {
	ruo.mutation.ResetSort()
	ruo.mutation.SetSort(u)
	return ruo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableSort(u *uint32) *RoleUpdateOne {
	if u != nil {
		ruo.SetSort(*u)
	}
	return ruo
}

// AddSort adds u to the "sort" field.
func (ruo *RoleUpdateOne) AddSort(u int32) *RoleUpdateOne {
	ruo.mutation.AddSort(u)
	return ruo
}

// AddMenuIDs adds the "menus" edge to the Menu entity by IDs.
func (ruo *RoleUpdateOne) AddMenuIDs(ids ...uint64) *RoleUpdateOne {
	ruo.mutation.AddMenuIDs(ids...)
	return ruo
}

// AddMenus adds the "menus" edges to the Menu entity.
func (ruo *RoleUpdateOne) AddMenus(m ...*Menu) *RoleUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ruo.AddMenuIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ruo *RoleUpdateOne) AddUserIDs(ids ...uuid.UUID) *RoleUpdateOne {
	ruo.mutation.AddUserIDs(ids...)
	return ruo
}

// AddUsers adds the "users" edges to the User entity.
func (ruo *RoleUpdateOne) AddUsers(u ...*User) *RoleUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.AddUserIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// ClearMenus clears all "menus" edges to the Menu entity.
func (ruo *RoleUpdateOne) ClearMenus() *RoleUpdateOne {
	ruo.mutation.ClearMenus()
	return ruo
}

// RemoveMenuIDs removes the "menus" edge to Menu entities by IDs.
func (ruo *RoleUpdateOne) RemoveMenuIDs(ids ...uint64) *RoleUpdateOne {
	ruo.mutation.RemoveMenuIDs(ids...)
	return ruo
}

// RemoveMenus removes "menus" edges to Menu entities.
func (ruo *RoleUpdateOne) RemoveMenus(m ...*Menu) *RoleUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ruo.RemoveMenuIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (ruo *RoleUpdateOne) ClearUsers() *RoleUpdateOne {
	ruo.mutation.ClearUsers()
	return ruo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (ruo *RoleUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *RoleUpdateOne {
	ruo.mutation.RemoveUserIDs(ids...)
	return ruo
}

// RemoveUsers removes "users" edges to User entities.
func (ruo *RoleUpdateOne) RemoveUsers(u ...*User) *RoleUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the RoleUpdate builder.
func (ruo *RoleUpdateOne) Where(ps ...predicate.Role) *RoleUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	ruo.defaults()
	return withHooks[*Role, RoleMutation](ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoleUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := role.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeUint64))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := ruo.mutation.AddedStatus(); ok {
		_spec.AddField(role.FieldStatus, field.TypeUint8, value)
	}
	if ruo.mutation.StatusCleared() {
		_spec.ClearField(role.FieldStatus, field.TypeUint8)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if value, ok := ruo.mutation.DefaultRouter(); ok {
		_spec.SetField(role.FieldDefaultRouter, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Sort(); ok {
		_spec.SetField(role.FieldSort, field.TypeUint32, value)
	}
	if value, ok := ruo.mutation.AddedSort(); ok {
		_spec.AddField(role.FieldSort, field.TypeUint32, value)
	}
	if ruo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: role.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedMenusIDs(); len(nodes) > 0 && !ruo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: role.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: role.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !ruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
