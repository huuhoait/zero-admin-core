// Code generated by ent, DO NOT EDIT.

package department

import (
	"time"
)

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAncestors holds the string denoting the ancestors field in the database.
	FieldAncestors = "ancestors"
	// FieldLeader holds the string denoting the leader field in the database.
	FieldLeader = "leader"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the department in the database.
	Table = "sys_departments"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "sys_departments"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "sys_departments"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_id"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "sys_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "sys_users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "department_id"
)

// Columns holds all SQL columns for department fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldSort,
	FieldName,
	FieldAncestors,
	FieldLeader,
	FieldPhone,
	FieldEmail,
	FieldRemark,
	FieldParentID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort uint32
	// DefaultParentID holds the default value on creation for the "parent_id" field.
	DefaultParentID uint64
)
