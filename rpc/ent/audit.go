// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/huuhoait/zero-admin-core/rpc/ent/audit"
)

// Audit is the model entity for the Audit schema.
type Audit struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Effect Object
	ObjectName string `json:"object_name,omitempty"`
	// Effect action
	ActionName string `json:"action_name,omitempty"`
	// Changed Data
	ChangedData string `json:"changed_data,omitempty"`
	// User Create
	CreatedBy string `json:"created_by,omitempty"`
	// User Update
	UpdatedBy string `json:"updated_by,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Audit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case audit.FieldID:
			values[i] = new(sql.NullInt64)
		case audit.FieldObjectName, audit.FieldActionName, audit.FieldChangedData, audit.FieldCreatedBy, audit.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case audit.FieldCreatedAt, audit.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Audit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Audit fields.
func (a *Audit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case audit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case audit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case audit.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case audit.FieldObjectName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field object_name", values[i])
			} else if value.Valid {
				a.ObjectName = value.String
			}
		case audit.FieldActionName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action_name", values[i])
			} else if value.Valid {
				a.ActionName = value.String
			}
		case audit.FieldChangedData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field changed_data", values[i])
			} else if value.Valid {
				a.ChangedData = value.String
			}
		case audit.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				a.CreatedBy = value.String
			}
		case audit.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				a.UpdatedBy = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Audit.
// Note that you need to call Audit.Unwrap() before calling this method if this Audit
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Audit) Update() *AuditUpdateOne {
	return NewAuditClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Audit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Audit) Unwrap() *Audit {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Audit is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Audit) String() string {
	var builder strings.Builder
	builder.WriteString("Audit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("object_name=")
	builder.WriteString(a.ObjectName)
	builder.WriteString(", ")
	builder.WriteString("action_name=")
	builder.WriteString(a.ActionName)
	builder.WriteString(", ")
	builder.WriteString("changed_data=")
	builder.WriteString(a.ChangedData)
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(a.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(a.UpdatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// Audits is a parsable slice of Audit.
type Audits []*Audit
